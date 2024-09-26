.PHONY: build
build:
	go build -v ./cmd/telegrambot

.DEFAULT_GOAL := build