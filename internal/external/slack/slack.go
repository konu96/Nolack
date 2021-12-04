package slack

import (
	"github.com/konu96/Nolack/internal/repository/dto"
	"github.com/slack-go/slack"
)

type Slack struct {
	Client *slack.Client
}

func NewSlack(client *slack.Client) *Slack {
	return &Slack{
		Client: client,
	}
}

func (s *Slack) Notify(input dto.NotifyInput) error {
	if _, _, err := s.Client.PostMessage(input.Channel, slack.MsgOptionText(input.Text, true)); err != nil {
		return err
	}
	return nil
}
