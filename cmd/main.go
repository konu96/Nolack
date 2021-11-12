package main

import (
	"github.com/konu96/Nolack/internal/external/server"
	"github.com/konu96/Nolack/internal/external/slack"
	"github.com/konu96/Nolack/internal/interfaces"
	slackGo "github.com/slack-go/slack"
	"log"
	"os"
)

func main() {
	controller := interfaces.NewController(slack.NewSlack(slackGo.New(os.Getenv("SLACK_BOT_TOKEN"))))
	s := server.NewServer(controller)

	if err := s.Run(); err != nil {
		log.Fatal(err)
	}
}
