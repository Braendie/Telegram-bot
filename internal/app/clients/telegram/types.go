package telegram

// Update represents a single update from the Telegram Bot API, typically containing
// information about incoming messages or other events
type Update struct {
	ID      int              `json:"update_id"`
	Message *IncomingMessage `json:"message"`
}

// UpdatesResponse is the structure for the response returned by the "getUpdates" API,
// containing a status and an array of updates
type UpdatesResponse struct {
	Ok     bool     `json:"ok"`
	Result []Update `json:"result"`
}

// IncomingMessage represents the structure of a received message in an update, including
// the message text, sender details, and chat information
type IncomingMessage struct {
	Text string `json:"text"`
	From From   `json:"from"`
	Chat Chat   `json:"chat"`
}

// From contains information about the user who sent the message
type From struct {
	UserName string `json:"username"`
}

// Chat contains details about the chat where the message was received
type Chat struct {
	ID int `json:"id"`
}
