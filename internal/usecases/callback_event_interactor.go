package usecases

import (
	"fmt"
	"github.com/konu96/Nolack/internal/domain/data"
	"github.com/konu96/Nolack/internal/external/slack"
	"github.com/konu96/Nolack/internal/usecases/dto"
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
		requestMessageCount := 4
		messages := strings.Split(event.Text, " ")
		if len(messages) < requestMessageCount {
			return &Error{
				StatusCode: http.StatusBadRequest,
				Err:        fmt.Errorf("missing argument: want %d argumets but got %d", requestMessageCount, len(message)),
			}
		}

		switch data.Command(messages[1]) {
		case data.Create:
			input := dto.PageInput{
				PageID: messages[2],
				URL:    "https://d3bhdfps5qyllw.cloudfront.net/org/63/63516e4f15e183b8925052964a58f077_1080x700_w.jpg",
			}
			if err := i.CreatePageInteractor.Exec(event.Channel, input); err != nil {
				return &Error{
					StatusCode: http.StatusInternalServerError,
					Err:        err,
				}
			}
		}
	}

	return nil
}
