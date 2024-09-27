package sqlstorage

import (
	"database/sql"
	"math/rand"

	"github.com/Braendie/Telegram-bot/internal/app/lib/e"
	"github.com/Braendie/Telegram-bot/internal/app/storage"
	_ "github.com/lib/pq"
)

// Storage implements Storage interface with PostgreSQL
type DBStorage struct {
	db *sql.DB
}

// New creates a new instance of Storage
func New(db *sql.DB) *DBStorage {
	return &DBStorage{
		db: db,
	}
}

// Save adds a new page to the database
func (s *DBStorage) Save(p *storage.Page) error {
	return s.db.QueryRow("INSERT INTO pages (username, url, tag, description) VALUES($1, $2, $3, $4) RETURNING id",
		p.UserName,
		p.URL,
		p.Tag,
		p.Description,
	).Scan(&p.ID)
}

// PickRandom retrieves a random page for a specific user
func (s *DBStorage) PickRandom(userName string) (*storage.Page, error) {
	pages := []*storage.Page{}
	query := `SELECT username, url, tag, description FROM pages WHERE username = $1`
	rows, err := s.db.Query(query, userName)
	if err != nil {
		return nil, e.Wrap("can't pick random page", err)
	}
	defer rows.Close()

	for rows.Next() {
		p := &storage.Page{}
		if err := rows.Scan(&p.UserName, &p.URL, &p.Tag, &p.Description); err != nil {
			return nil, err
		}
		pages = append(pages, p)
	}

	if len(pages) == 0 {
		return nil, storage.ErrNoSavedPages
	}

	// Select a random page
	randompage := pages[rand.Intn(len(pages))]
	return &storage.Page{
		UserName:    randompage.UserName,
		URL:         randompage.URL,
		Tag:         randompage.Tag,
		Description: randompage.Description,
	}, nil
}

// Remove a page from the database
func (s *DBStorage) Remove(p *storage.Page) error {
	query := `DELETE FROM pages WHERE url = $1 AND username = $2`
	_, err := s.db.Exec(query, p.URL, p.UserName)
	if err != nil {
		return e.Wrap("can't remove page", err)
	}
	return nil
}

// IsExists checks if a page already exists for a specific user
func (s *DBStorage) IsExists(p *storage.Page) (bool, error) {
	query := `SELECT COUNT(*) FROM pages WHERE url = $1 AND username = $2`
	var count int
	if err := s.db.QueryRow(query, p.URL, p.UserName).Scan(&count); err != nil {
		return false, e.Wrap("can't check if page exists", err)
	}
	return count > 0, nil
}
