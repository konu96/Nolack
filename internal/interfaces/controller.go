package interfaces

import (
	"fmt"
	"github.com/konu96/Nolack/internal/external/notion"
	"github.com/konu96/Nolack/internal/external/slack"
	"github.com/konu96/Nolack/internal/repository"
	"github.com/konu96/Nolack/internal/usecases"
	"github.com/slack-go/slack/slackevents"
	"io/ioutil"
	"net/http"
)

type Controller struct {
	VerifyURLUseCase     usecases.VerifyURLUseCase
	CallbackEventUseCase usecases.CallbackEventUseCase
}

func NewController(slack *slack.Slack) Controller {
	return Controller{
		VerifyURLUseCase: usecases.NewVerifyURLInteractor(),
		CallbackEventUseCase: usecases.NewCallbackEventInteractor(
			repository.NewNotionRepository(notion.NewNotion()),
			usecases.NewNotifyInteractor(slack),
		),
	}
}

func (c *Controller) Exec(w http.ResponseWriter, r *http.Request) error {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return err
	}

	event, err := slackevents.ParseEvent(body, slackevents.OptionNoVerifyToken())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return err
	}

	switch event.Type {
	case slackevents.URLVerification:
		if err := c.VerifyURLUseCase.Exec(w, body); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return fmt.Errorf("failed to url verification: %w", err)
		}

	case slackevents.CallbackEvent:
		if err := c.CallbackEventUseCase.Exec(event); err != nil {
			w.WriteHeader(err.StatusCode)
			return fmt.Errorf("failed to callback event: %w", err.Err)
		}
	}

	return nil
}
