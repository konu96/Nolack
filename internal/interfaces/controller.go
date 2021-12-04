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
	VerifyURLInteractor     usecases.VerifyURLInteractor
	CallbackEventInteractor usecases.CallbackEventInteractor
}

func NewController(slack *slack.Slack) Controller {
	return Controller{
		VerifyURLInteractor:     usecases.NewVerifyURLInteractor(),
		CallbackEventInteractor: usecases.NewCallbackEventInteractor(slack, repository.NewNotionRepository(notion.NewNotion())),
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
		if err := c.VerifyURLInteractor.Exec(w, body); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return fmt.Errorf("url verification: %w", err)
		}

	case slackevents.CallbackEvent:
		if err := c.CallbackEventInteractor.Exec(event); err != nil {
			w.WriteHeader(err.StatusCode)
			return fmt.Errorf("callback event: %w", err.Err)
		}
	}

	return nil
}
