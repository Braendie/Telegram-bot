package main

import (
	"database/sql"
	"flag"
	"log"

	"github.com/Braendie/Telegram-bot/config"
	tgClient "github.com/Braendie/Telegram-bot/internal/app/clients/telegram"
	event_consumer "github.com/Braendie/Telegram-bot/internal/app/consumer/event-consumer"
	"github.com/Braendie/Telegram-bot/internal/app/events/telegram"
	"github.com/Braendie/Telegram-bot/internal/app/storage/sqlstorage"
	"github.com/BurntSushi/toml"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "config/telegram-bot.toml", "path to config file")
}

func main() {
	flag.Parse()
	config := config.NewConfig()
	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		log.Fatal("can't decode toml file", err)
	}

	tgClient := tgClient.New(config.TGBotHost, config.TGToken)

	db, err := newDB(config.DatabaseURL)
	if err != nil {
		log.Fatal("can't create database", err)
	}

	defer db.Close()
	storage := sqlstorage.New(db)
	eventsProcessor := telegram.New(tgClient, storage)

	log.Print("service started")
	consumer := event_consumer.New(eventsProcessor, eventsProcessor, config.BatchSize)

	if err := consumer.Start(); err != nil {
		log.Fatal("service is stopped", err)
	}
}

func newDB(databaseURL string) (*sql.DB, error) {
	db, err := sql.Open("postgres", databaseURL)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
