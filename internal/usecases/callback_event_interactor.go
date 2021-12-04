package usecases

import (
	"fmt"
	"github.com/konu96/Nolack/internal/domain/data"
	"github.com/konu96/Nolack/internal/external/slack"
	"github.com/konu96/Nolack/internal/usecases/repository"
	"github.com/slack-go/slack/slackevents"
	"net/http"
	"strings"
)

type CallbackEventInteractor struct {
	CreatePageInteractor CreateNotionPageInteractor
}

func NewCallbackEventInteractor(slack *slack.Slack, notionRepository repository.NotionRepository) CallbackEventInteractor {
	return CallbackEventInteractor{
		CreatePageInteractor: NewCreatePageInteractor(slack, notionRepository),
	}
}

func (i *CallbackEventInteractor) Exec(w http.ResponseWriter, event slackevents.EventsAPIEvent) error {
	innerEvent := event.InnerEvent

	switch event := innerEvent.Data.(type) {
	case *slackevents.AppMentionEvent:
		message := strings.Split(event.Text, " ")
		if len(message) < 2 {
			w.WriteHeader(http.StatusBadRequest)
			return fmt.Errorf("missing argument: want 2 argumets but got %d", len(message))
		}

		switch data.Command(message[1]) {
		case data.Create:
			if err := i.CreatePageInteractor.Exec(event.Channel); err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return err
			}
		}
	}

	return nil
}
