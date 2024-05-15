package main

import (
	"encoding/json"
	"net/http"

	log "github.com/sirupsen/logrus"

	"github.com/wojtekwanczyk/github-review-bot/pkg/config"
	"github.com/wojtekwanczyk/github-review-bot/pkg/github"
	"github.com/wojtekwanczyk/github-review-bot/pkg/webex"
)

func main() {
	http.HandleFunc("/", Handler)
	log.Fatal(http.ListenAndServe(":7001", nil))

}

func Handler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var hookData github.Hook
	err := decoder.Decode(&hookData)
	if err != nil {
		log.Errorf("Cannot decode request: %s", err)
	}

	log.Infof("Request data: %+v", hookData)

	for _, user := range hookData.PullRequest.RequestedReviewers {
		if roomID, ok := config.RoomMapping[user.Login]; ok {
			webex.SendMessage("Test message", roomID)
		}
	}
}
