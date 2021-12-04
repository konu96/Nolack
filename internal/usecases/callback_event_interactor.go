package usecases

import (
	"fmt"
	"github.com/konu96/Nolack/internal/domain/data"
	repositorydto "github.com/konu96/Nolack/internal/repository/dto"
	usecasesdto "github.com/konu96/Nolack/internal/usecases/dto"
	"github.com/konu96/Nolack/internal/usecases/repository"
	"github.com/slack-go/slack/slackevents"
	"net/http"
	"strings"
)

type CallbackEventUseCase interface {
	Exec(event slackevents.EventsAPIEvent) *Error
}

type CallbackEventInteractor struct {
	CreateNotionPageUseCase CreateNotionPageUseCase
	NotifyUseCase           NotifyUseCase
}

func NewCallbackEventInteractor(
	notionRepository repository.NotionRepository,
	notifyInteractor *NotifyInteractor,
) *CallbackEventInteractor {
	return &CallbackEventInteractor{
		CreateNotionPageUseCase: NewCreatePageInteractor(notionRepository, notifyInteractor),
		NotifyUseCase:           notifyInteractor,
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
		requestMessageCount := 3

		messages := strings.Split(event.Text, " ")
		if len(messages) < requestMessageCount {
			return &Error{
				StatusCode: http.StatusBadRequest,
				Err:        fmt.Errorf("missing argument: want %d argumets but got %d", requestMessageCount, len(messages)),
			}
		}

		switch data.Command(messages[1]) {
		case data.CreatePage:
			input := usecasesdto.CreatePageInput{
				PageID: messages[2],
				URL:    "https://d3bhdfps5qyllw.cloudfront.net/org/63/63516e4f15e183b8925052964a58f077_1080x700_w.jpg",
			}
			if err := i.CreateNotionPageUseCase.Exec(event.Channel, input); err != nil {
				err := i.NotifyUseCase.Exec(repositorydto.NotifyInput{
					Channel: event.Channel,
					Text:    err.Error(),
				})
				return &Error{
					StatusCode: http.StatusInternalServerError,
					Err:        err,
				}
			}
		case data.CreateDatabase:
			break
		}
	}

	return nil
}
