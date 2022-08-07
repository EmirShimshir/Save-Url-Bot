package main

import (
	"flag"
	"log"

	tgClient "Save_Url_Bot/client/telegram"
	event_consumer "Save_Url_Bot/consumer/event-consumer"
	"Save_Url_Bot/events/telegram"
	"Save_Url_Bot/storage/files"
)

const (
	tgBotHost   = "api.telegram.org"
	storagePath = "storage/files"
	batchSize   = 100
)

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
	token := flag.String(
		"tg-bot-token",
		"",
		"token for access to telegram bot",
	)

	flag.Parse()

	if *token == "" {
		log.Fatal("token is not specified")
	}

	return *token
}
