package main

import (
	"flag"
	"log"
	tgClient "reminder-bot/clients/telegram"
	event_consumer "reminder-bot/consumer/event-consumer"
	"reminder-bot/events/telegram"
	"reminder-bot/storage/files"
)

const (
	tgBotHost   = "api.telegram.org"
	storagePath = "files_storage"
	batchSize   = 100
)

/*
			TOKEN:
	7809238484:AAE0v-t6q0kqpjHbIcEEWggs9C3t58pyvA4

			TO RUN BOT:
	go build
	.\reminder-bot -tg-bot-token '7809238484:AAE0v-t6q0kqpjHbIcEEWggs9C3t58pyvA4'
*/

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
