package events

// Fetcher defines an interface for fetching events, specifying a Fetch method
// that retrieves a batch of events up to a specified limit
type Fetcher interface {
	Fetch(limit int) ([]Event, error)
}

// Processor defines an interface for processing individual events, with a Process
// method that takes an Event and returns an error if processing fails
type Processor interface {
	Process(e Event) error
}

type Type int

const (
	Unknown Type = iota
	Message
)

// Event defines the structure of an event, including its type, content text, and
// optional metadata
type Event struct {
	Type Type
	Text string
	Meta interface{}
}
