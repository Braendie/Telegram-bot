package telegram

import (
	"errors"

	"github.com/Braendie/Telegram-bot/internal/app/clients/telegram"
	"github.com/Braendie/Telegram-bot/internal/app/events"
	"github.com/Braendie/Telegram-bot/internal/app/lib/e"
	"github.com/Braendie/Telegram-bot/internal/app/storage"
)

// Processor handles event processing from the Telegram client and manages interactions with the storage
type Processor struct {
	tg      *telegram.Client
	offset  int
	storage storage.Storage
}

// Meta contains metadata about a Telegram message, including chat ID and username
type Meta struct {
	ChatID   int
	UserName string
}

var (
	ErrUnknownEvent    = errors.New("unknown event type")
	ErrUnknownMetaType = errors.New("unknown meta type")
)

// New creates and returns a new Processor with the provided Telegram client and storage
func New(client *telegram.Client, storage storage.Storage) *Processor {
	return &Processor{
		tg:      client,
		storage: storage,
	}
}

// Fetch retrieves a batch of events from the Telegram API and returns them as an array of Event structs
func (p *Processor) Fetch(limit int) ([]events.Event, error) {
	updates, err := p.tg.Updates(p.offset, limit)
	if err != nil {
		return nil, e.Wrap("can't get events", err)
	}

	if len(updates) == 0 {
		return nil, nil
	}
	res := make([]events.Event, 0, len(updates))

	for _, u := range updates {
		res = append(res, event(u))
	}

	p.offset = updates[len(updates)-1].ID + 1

	return res, nil
}

// Process takes an Event and processes it based on its type
func (p *Processor) Process(event events.Event) error {
	switch event.Type {
	case events.Message:
		return p.processMessage(event)
	default:
		return e.Wrap("can't process message", ErrUnknownEvent)
	}
}

// processMessage handles a message event, retrieving metadata and executing commands
func (p *Processor) processMessage(event events.Event) error {
	meta, err := meta(event)
	if err != nil {
		return e.Wrap("can't process message", err)
	}

	if err := p.doCmd(event.Text, meta.ChatID, meta.UserName); err != nil {
		return e.Wrap("can't process messsage", err)
	}

	return nil
}

// meta extracts and returns metadata from an event, ensuring it matches the expected Meta type
func meta(event events.Event) (Meta, error) {
	res, ok := event.Meta.(Meta)
	if !ok {
		return Meta{}, e.Wrap("can't get meta", ErrUnknownMetaType)
	}

	return res, nil
}

// event converts a Telegram Update into an Event, extracting relevant information
func event(upd telegram.Update) events.Event {
	updType := fetchType(upd)

	res := events.Event{
		Type: updType,
		Text: fetchText(upd),
	}

	if updType == events.Message {
		res.Meta = Meta{
			ChatID:   upd.Message.Chat.ID,
			UserName: upd.Message.From.UserName,
		}
	}

	return res
}

// fetchType determines the event type based on the contents of the Telegram Update
func fetchType(upd telegram.Update) events.Type {
	if upd.Message == nil {
		return events.Unknown
	}

	return events.Message
}

// fetchText retrieves the text content from the Telegram Update, if available
func fetchText(upd telegram.Update) string {
	if upd.Message == nil {
		return ""
	}

	return upd.Message.Text
}
