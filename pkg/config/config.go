package config

import (
	"encoding/json"
	"os"

	"github.com/kelseyhightower/envconfig"
	log "github.com/sirupsen/logrus"
)

type config struct {
	BotToken        string `split_words:"true" required:"true"`
	RoomMappingPath string `split_words:"true" required:"true"`
}

var Webex config

func init() {
	err := envconfig.Process("webex", &Webex)
	if err != nil {
		log.Fatalf("Cannot parse config: %s", err)
	}

	readMapping(Webex.RoomMappingPath)
}

type roomMapping map[string]string

var RoomMapping roomMapping

func readMapping(filepath string) {
	mappingFile, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}
	defer mappingFile.Close()

	var decoder *json.Decoder = json.NewDecoder(mappingFile)
	if err != nil {
		log.Fatal(err)
	}

	err = decoder.Decode(&RoomMapping)
	if err != nil {
		log.Fatal(err)
	}

	log.Infof("Room Mapping: %s", RoomMapping)
}
