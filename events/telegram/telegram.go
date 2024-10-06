package telegram

import (
	"reminder-bot/clients/telegram"

	"golang.org/x/mod/sumdb/storage"
)

type Processor struct {
	tg     *telegram.Client
	offset int
	// storage
}

func New(client &telegram.Client)
