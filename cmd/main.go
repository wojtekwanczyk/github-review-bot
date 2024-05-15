package main

import (
	webex "github.com/jbogarin/go-cisco-webex-teams/sdk"
	"github.com/kelseyhightower/envconfig"
	log "github.com/sirupsen/logrus"
)

type Config struct {
	BotToken string `split_words:"true" required:"true"`
	RoomID   string `split_words:"true" required:"true"`
}

var config Config

func parseConfig() {
	err := envconfig.Process("webex", &config)
	if err != nil {
		log.Fatalf("Cannot parse config: %s", err)
	}
}

func main() {
	parseConfig()

	c := webex.NewClient()
	c.SetAuthToken(config.BotToken)

	msgRequest := &webex.MessageCreateRequest{
		Text:   "Test message",
		RoomID: config.RoomID,
	}

	log.Debugf("Message request: %+v", *msgRequest)

	msg, _, err := c.Messages.CreateMessage(msgRequest)
	if err != nil {
		log.Fatalf("Cannot send webex message: %s", err)
	}

	log.Infof("Message sent: %+v", *msg)
}
