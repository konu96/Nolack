package slack

import (
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
