package telegram

import (
	"database/sql"
	"errors"
	"log"
	"net/url"
	"regexp"
	"strings"

	"github.com/Braendie/Telegram-bot/internal/app/lib/e"
	"github.com/Braendie/Telegram-bot/internal/app/storage"
)

const (
	RndCmd    = "/rnd"
	HelpCmdEn = "/help_en"
	HelpCmdRu = "/help_ru"
	StartCmd  = "/start"
)

func (p *Processor) doCmd(text string, chatID int, username string) error {
	text = strings.TrimSpace(text)

	log.Printf("got new command '%s' from '%s'", text, username)

	words := strings.Split(text, " ")

	if isAddCmd(words[0]) {
		if len(words) < 2 {
			return p.savePage(chatID, words[0], username, "", "")
		} else if len(words) < 3 {
			return p.savePage(chatID, words[0], username, words[1], "")
		}
		return p.savePage(chatID, words[0], username, words[1], strings.Join(words[2:], " "))
	}
	if len(text) != 0 {
		if text[:1] != "/" {
			return p.tg.SendMessage(chatID, "okay.")
		}
	}

	if strings.HasPrefix(text, RndCmd) {
		return p.sendRandom(chatID, username)
	} else if strings.HasPrefix(text, HelpCmdEn) {
		return p.sendHelpEn(chatID)
	} else if strings.HasPrefix(text, StartCmd) {
		return p.sendHello(chatID)
	} else if strings.HasPrefix(text, HelpCmdRu) {
		return p.sendHelpRu(chatID)
	} else {
		return p.tg.SendMessage(chatID, msgUnknownCommand)
	}
}

func (p *Processor) savePage(chatID int, pageURL, username, tag, description string) error {
	page := &storage.Page{
		URL:         pageURL,
		UserName:    username,
		Tag:         sql.NullString{String: tag, Valid: tag != ""},
		Description: sql.NullString{String: description, Valid: description != ""},
	}

	IsExists, err := p.storage.IsExists(page)
	if err != nil {
		return e.Wrap("can't do command: save page", err)
	}
	if IsExists {
		return p.tg.SendMessage(chatID, msgAlreadyExists)
	}

	if err := p.storage.Save(page); err != nil {
		return e.Wrap("can't do command: save page", err)
	}

	if err := p.tg.SendMessage(chatID, msgSaved); err != nil {
		return e.Wrap("can't do command: save page", err)
	}

	return nil
}

func (p *Processor) sendRandom(chatID int, username string) error {
	page, err := p.storage.PickRandom(username)
	if err != nil && !errors.Is(err, storage.ErrNoSavedPages) {
		return e.Wrap("can't do command: send random", err)
	}

	if errors.Is(err, storage.ErrNoSavedPages) {
		return p.tg.SendMessage(chatID, msgNoSavedPages)
	}

	if page.Tag.Valid && !strings.HasPrefix("#", page.Tag.String) {
		page.Tag.String = "#" + page.Tag.String
	}

	if err := p.tg.SendMessage(chatID, page.URL+"\n"+page.Tag.String+"\n"+page.Description.String); err != nil {
		return e.Wrap("can't do command: send random", err)
	}

	return p.storage.Remove(page)
}

func (p *Processor) sendHelpEn(chatID int) error {
	return p.tg.SendMessage(chatID, msgHelpEn+msgHelpCmdEn)
}

func (p *Processor) sendHelpRu(chatID int) error {
	return p.tg.SendMessage(chatID, msgHelpRu+msgHelpCmdRu)
}

func (p *Processor) sendHello(chatID int) error {
	return p.tg.SendMessage(chatID, msgHelloEn)
}

func isAddCmd(text string) bool {
	return isURL(text)
}

func isURL(text string) bool {
	if !strings.HasPrefix(text, "http://") && !strings.HasPrefix(text, "https://") {
		text = "http://" + text
	}

	u, err := url.Parse(text)
	if err != nil || u.Host == "" {
		return false
	}

	// Проверка хоста на валидность (домен или IP)
	hostname := u.Hostname()
	// Регулярное выражение для проверки корректности доменного имени или IP-адреса
	validHost := regexp.MustCompile(`^([a-zA-Z0-9-]+\.)+[a-zA-Z]{2,6}$|^(\d{1,3}\.){3}\d{1,3}$`)
	return validHost.MatchString(hostname)
}
