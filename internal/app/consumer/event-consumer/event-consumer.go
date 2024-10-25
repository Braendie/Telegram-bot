package event_consumer

import (
	"log"
	"time"

	"github.com/Braendie/Telegram-bot/internal/app/events"
)

// Consumer represents an event consumer that fetches and processes events in batches
type Consumer struct {
	fetcher   events.Fetcher
	processor events.Processor
	batchSize int
}

// New creates and returns a new Consumer with the given fetcher, processor, and batch size
func New(fetcher events.Fetcher, processor events.Processor, batchSize int) Consumer {
	return Consumer{
		fetcher:   fetcher,
		processor: processor,
		batchSize: batchSize,
	}
}

// Start initiates the event consumption loop, fetching and processing events in batches
func (c Consumer) Start() error {
	for {
		gotEvents, err := c.fetcher.Fetch(c.batchSize)
		if err != nil {
			log.Printf("[ERR] consumer: %s", err.Error())

			continue
		}

		if len(gotEvents) == 0 {
			time.Sleep(time.Second)

			continue
		}

		if err := c.handleEvents(gotEvents); err != nil {
			log.Print(err)

			continue
		}
	}
}

// handleEvents processes each event in the provided slice of events
func (c *Consumer) handleEvents(events []events.Event) error {
	for _, event := range events {
		log.Printf("got new event: %s", event.Text)

		if err := c.processor.Process(event); err != nil {
			log.Printf("can't handle event: %s", err.Error())

			continue
		}
	}

	return nil
}
