package telegram

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"path"
	"strconv"

	"github.com/Braendie/Telegram-bot/internal/app/lib/e"
)

// Client represents a Telegram client that communicates with the Telegram Bot API
type Client struct {
	host     string
	basePath string
	client   http.Client
}

// New initializes and returns a new Client with the provided host and token
func New(host, token string) *Client {
	return &Client{
		host:     host,
		basePath: newBasePath(token),
		client:   http.Client{},
	}
}

// newBasePath constructs the base API path using the bot token
func newBasePath(token string) string {
	return "bot" + token
}

// Updates retrieves new updates (messages, commands, etc.) from the Telegram API with the specified offset and limit
func (c *Client) Updates(offset, limit int) ([]Update, error) {
	q := url.Values{}
	q.Add("offset", strconv.Itoa(offset))
	q.Add("limit", strconv.Itoa(limit))

	data, err := c.doRequest("getUpdates", q)
	if err != nil {
		return nil, e.Wrap("can't get updates", err)
	}

	var res UpdatesResponse

	if err := json.Unmarshal(data, &res); err != nil {
		return nil, e.Wrap("can't get updates", err)
	}

	return res.Result, nil
}

// SendMessage sends a text message to a specified chat ID
func (c *Client) SendMessage(chatID int, text string) error {
	q := url.Values{}
	q.Add("chat_id", strconv.Itoa(chatID))
	q.Add("text", text)

	_, err := c.doRequest("sendMessage", q)
	if err != nil {
		return e.Wrap("can't send message", err)
	}

	return nil
}

// doRequest sends an HTTP GET request to the specified Telegram API method with query parameters
// It returns the response body as a byte slice
func (c *Client) doRequest(method string, query url.Values) ([]byte, error) {
	u := url.URL{
		Scheme: "https",
		Host:   c.host,
		Path:   path.Join(c.basePath, method),
	}

	req, err := http.NewRequest(http.MethodGet, u.String(), nil)
	if err != nil {
		return nil, e.Wrap("can't do request", err)
	}

	req.URL.RawQuery = query.Encode()

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, e.Wrap("can't do request", err)
	}

	defer func() { _ = resp.Body.Close() }()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, e.Wrap("can't do request", err)
	}

	return body, nil
}
