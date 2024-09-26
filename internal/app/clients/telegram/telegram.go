package telegram

import (
	"net/http"

	"github.com/Braendie/Telegram-bot/internal/app/clients"
)

type Client struct {
	host     string
	basePath string
	client   http.Client
}

func New(host, token string) Client {
	return Client{
		host:     host,
		basePath: "bot" + token,
		client:   http.Client{},
	}
}

func (c *Client) Updates() ([]clients.Update, error) {

}

func (c *Client) SendMessage() {

}
