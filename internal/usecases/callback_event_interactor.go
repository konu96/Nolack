package usecases

import (
	"github.com/konu96/Nolack/internal/external/slack"
	slackGo "github.com/slack-go/slack"
	"github.com/slack-go/slack/slackevents"
	"log"
	"net/http"
	"os"
	"strings"
)

type CallbackEventInteractor struct {
}

func NewCallbackEventInteractor() CallbackEventInteractor {
	return CallbackEventInteractor{}
}

func (i *CallbackEventInteractor) Exec(w http.ResponseWriter, event slackevents.EventsAPIEvent) {
	innerEvent := event.InnerEvent

	switch event := innerEvent.Data.(type) {
	case *slackevents.AppMentionEvent:
		message := strings.Split(event.Text, " ")
		if len(message) < 2 {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		command := message[1]
		user := event.User

		sl := slack.NewSlack(slackGo.New(os.Getenv("SLACK_BOT_TOKEN")))
		switch command {
		case "hello":
			if _, _, err := sl.Client.PostMessage(event.Channel, slackGo.MsgOptionText("<@"+user+"> world", false)); err != nil {
				log.Println(err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
		}
	}
}
