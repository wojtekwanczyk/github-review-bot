package main

import (
	"io"
	"net/http"

	log "github.com/sirupsen/logrus"
)

func main() {
	http.HandleFunc("/", Handler)
	log.Fatal(http.ListenAndServe(":7001", nil))

	// webex.SendMessage("Test message", config.Webex.RoomID)
}

type GithubData struct {
}

func Handler(w http.ResponseWriter, r *http.Request) {
	// decoder := json.NewDecoder(r.Body)
	// var gd GithubData
	// err := decoder.Decode(&gd)
	// if err != nil {
	// 	log.Errorf("Cannot decode request: %s", err)
	// }

	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Errorf("%s", err)
	}
	log.Info("Request data: %s", body)
}
