package main

import (
	"log"

	tgClient "Save_Url_Bot/client/telegram"
	event_consumer "Save_Url_Bot/consumer/event-consumer"
	"Save_Url_Bot/events/telegram"
	"Save_Url_Bot/storage/files"

	"github.com/kelseyhightower/envconfig"
)

const (
	tgBotHost   = "api.telegram.org"
	storagePath = "storage/files"
	batchSize   = 100
)

type Config struct {
	Token string
}

func main() {

	eventsProcessor := telegram.New(
		tgClient.New(tgBotHost, mustToken()),
		files.New(storagePath),
	)

	log.Print("service started")

	consumer := event_consumer.New(eventsProcessor, eventsProcessor, batchSize)

	if err := consumer.Start(); err != nil {
		log.Fatal("service is stopped", err)
	}
}

func mustToken() string {
	var cfg Config

	if err := envconfig.Process("tg", &cfg); err != nil {
		log.Fatal("token is not specified")
	}

	return cfg.Token
}
