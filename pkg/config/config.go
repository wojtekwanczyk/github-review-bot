package config

import (
	"github.com/kelseyhightower/envconfig"
	log "github.com/sirupsen/logrus"
)

type Config struct {
	BotToken string `split_words:"true" required:"true"`
	RoomID   string `split_words:"true" required:"true"`
}

var Webex Config

func init() {
	err := envconfig.Process("webex", &Webex)
	if err != nil {
		log.Fatalf("Cannot parse config: %s", err)
	}
}
