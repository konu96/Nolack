package usecases

import (
	"github.com/konu96/Nolack/internal/external/slack"
	slackGo "github.com/slack-go/slack"
	"log"
	"net/http"
)

type CreatePageInteractor struct {
	slack.Slack
}

func NewCreatePageInteractor(slack slack.Slack) CreatePageInteractor {
	return CreatePageInteractor{
		slack,
	}
}

func (i *CreatePageInteractor) Exec(w http.ResponseWriter, channel string) {
	if _, _, err := i.Client.PostMessage(channel, slackGo.MsgOptionText("a", false)); err != nil {
		log.Printf("can not post message: %s", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
