package main

import (
	"github.com/wojtekwanczyk/github-review-bot/pkg/config"
	"github.com/wojtekwanczyk/github-review-bot/pkg/webex"
)

func main() {
	webex.SendMessage("Test message", config.Webex.RoomID)
}
