package usecases

import (
	"github.com/konu96/Nolack/internal/domain/data"
	"github.com/konu96/Nolack/internal/external/slack"
	"github.com/slack-go/slack/slackevents"
	"net/http"
	"strings"
)

type CallbackEventInteractor struct {
	CreatePageInteractor CreatePageInteractor
}

func NewCallbackEventInteractor(slack slack.Slack, notion NotionInterface) CallbackEventInteractor {
	return CallbackEventInteractor{
		CreatePageInteractor: CreatePageInteractor{
			Slack:  slack,
			Notion: notion,
		},
	}
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

		c := data.Command(message[1])

		switch c {
		case data.Create:
			if err := i.CreatePageInteractor.Exec(event.Channel); err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
		}
	}
}
