package main

import (
	"flag"
	"log"

	tgClient "github.com/Braendie/Telegram-bot/internal/app/clients/telegram"
	event_consumer "github.com/Braendie/Telegram-bot/internal/app/consumer/event-consumer"
	"github.com/Braendie/Telegram-bot/internal/app/events/telegram"
	"github.com/Braendie/Telegram-bot/internal/app/storage/files"
)

const (
	tgBotHost   = "api.telegram.org"
	storagePath = "storage"
	batchSize   = 100
)

func main() {
	t := mustToken()

	tgClient := tgClient.New(tgBotHost, t)

	eventsProcessor := telegram.New(tgClient, files.New(storagePath))

	log.Print("service started")
	consumer := event_consumer.New(eventsProcessor, eventsProcessor, batchSize)

	if err := consumer.Start(); err != nil {
		log.Fatal("service is stopped", err)
	}
}

func mustToken() string {
	token := flag.String("tg-bot-token", "", "token for access to telegram bot")

	flag.Parse()

	if *token == "" {
		log.Fatal("token is not specified")
	}

	return *token
}
