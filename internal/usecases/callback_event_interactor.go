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

type Error struct {
	StatusCode int
	Err        error
}

func (i *CallbackEventInteractor) Exec(event slackevents.EventsAPIEvent) *Error {
	innerEvent := event.InnerEvent

	switch event := innerEvent.Data.(type) {
	case *slackevents.AppMentionEvent:
		message := strings.Split(event.Text, " ")
		if len(message) < 2 {
			return &Error{
				StatusCode: http.StatusBadRequest,
				Err:        fmt.Errorf("missing argument: want 2 argumets but got %d", len(message)),
			}
		}

		switch data.Command(message[1]) {
		case data.Create:
			if err := i.CreatePageInteractor.Exec(event.Channel); err != nil {
				return &Error{
					StatusCode: http.StatusInternalServerError,
					Err:        err,
				}
			}
		}
	}

	return nil
}
