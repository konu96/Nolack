package interfaces

import (
	"encoding/json"
	"github.com/konu96/Nolack/internal/external/slack"
	"github.com/konu96/Nolack/internal/usecases"
	"github.com/slack-go/slack/slackevents"
	"io/ioutil"
	"log"
	"net/http"
)

type Controller struct {
	VerifyURLInteractor     usecases.VerifyURLInteractor
	CallbackEventInteractor usecases.CallbackEventInteractor
}

func NewController(slack slack.Slack) Controller {
	return Controller{
		VerifyURLInteractor: usecases.NewVerifyURLInteractor(),
		CallbackEventInteractor: usecases.NewCallbackEventInteractor(slack),
	}
}

func (c *Controller) Exec(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	event, err := slackevents.ParseEvent(json.RawMessage(body), slackevents.OptionNoVerifyToken())
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	switch event.Type {
	case slackevents.URLVerification:
		c.VerifyURLInteractor.Exec(w, body)

	case slackevents.CallbackEvent:
		c.CallbackEventInteractor.Exec(w, event)
	}
}
