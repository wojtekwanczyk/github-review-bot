package main

import (
	"encoding/json"
	"net/http"

	log "github.com/sirupsen/logrus"

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

	log.Infof("Hook Data: %+v", hookData)

	// TODO: This logic must be rechecked in later phase
	if hookData.Action == "submitted" && hookData.Review.ID != 0 {
		webex.NotifyComment(&hookData)
	} else if hookData.Action == "opened" {
		webex.NotifyNewPr(&hookData.PullRequest)
	} else if hookData.Action == "synchronize" {
		webex.NotifyPrUpdated(&hookData.PullRequest)
	} else {
		log.Errorf("Unknown action: %s", hookData.Action)
	}

}
