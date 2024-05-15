package webex

import (
	"fmt"

	webex "github.com/jbogarin/go-cisco-webex-teams/sdk"
	log "github.com/sirupsen/logrus"

	"github.com/wojtekwanczyk/github-review-bot/pkg/config"
)

var client *webex.Client

func init() {
	client = webex.NewClient()
	client.SetAuthToken(config.Webex.BotToken)
}

func SendMessage(text, RoomID string) error {
	msgRequest := &webex.MessageCreateRequest{
		Markdown: text,
		RoomID:   RoomID,
	}

	log.Debugf("Message request: %+v", *msgRequest)

	msg, _, err := client.Messages.CreateMessage(msgRequest)
	if err != nil {
		return fmt.Errorf("cannot send webex message: %s", err)
	}

	log.Infof("Message sent: %+v", *msg)
	return nil
}
