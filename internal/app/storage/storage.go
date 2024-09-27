package storage

import (
	"crypto/sha1"
	"database/sql"
	"errors"
	"fmt"
	"io"

	"github.com/Braendie/Telegram-bot/internal/app/lib/e"
)

var ErrNoSavedPages = errors.New("no saved page")

type Storage interface {
	Save(p *Page) error
	PickRandom(userName string) (*Page, error)
	Remove(p *Page) error
	IsExists(p *Page) (bool, error)
}

// Page represents the page structure
type Page struct {
	ID          int
	URL         string
	UserName    string
	Tag         sql.NullString
	Description sql.NullString
}

func (p Page) Hash() (string, error) {
	h := sha1.New()

	if _, err := io.WriteString(h, p.URL); err != nil {
		return "", e.Wrap("can't calculate hash", err)
	}

	if _, err := io.WriteString(h, p.UserName); err != nil {
		return "", e.Wrap("can't calculate hash", err)
	}

	return fmt.Sprintf("%x", h.Sum(nil)), nil
}
