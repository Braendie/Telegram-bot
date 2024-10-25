package files

import (
	"encoding/gob"
	"errors"
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"time"

	"github.com/Braendie/Telegram-bot/internal/app/lib/e"
	"github.com/Braendie/Telegram-bot/internal/app/storage"
)

const defaultPerm = 0774

// Storage handles file-based storage for user pages
type Storage struct {
	basePath string
}

// New initializes and returns a new Storage instance with the specified base path
func New(basePath string) Storage {
	return Storage{
		basePath: basePath,
	}
}

// Save stores a page for a user by encoding it and saving it to a file
func (s Storage) Save(page *storage.Page) error {
	fPath := filepath.Join(s.basePath, page.UserName)

	if err := os.MkdirAll(fPath, defaultPerm); err != nil {
		return e.Wrap("can't save page", err)
	}

	fName, err := fileName(page)
	if err != nil {
		return e.Wrap("can't save page", err)
	}

	fPath = filepath.Join(fPath, fName)
	file, err := os.Create(fPath)
	if err != nil {
		return e.Wrap("can't save page", err)
	}
	defer func() { _ = file.Close() }()

	if err := gob.NewEncoder(file).Encode(page); err != nil {
		return e.Wrap("can't save page", err)
	}

	return nil
}

// PickRandom selects a random page from the user's stored pages
func (s Storage) PickRandom(userName string) (*storage.Page, error) {
	path := filepath.Join(s.basePath, userName)

	files, err := os.ReadDir(path)
	if err != nil {
		return nil, e.Wrap("can't pick random page", err)
	}

	if len(files) == 0 {
		return nil, storage.ErrNoSavedPages
	}

	src := rand.NewSource(time.Now().UnixNano())
	r := rand.New(src)
	n := r.Intn(len(files))

	file := files[n]

	return s.decodePage(filepath.Join(path, file.Name()))
}

// Remove deletes the specified page file from the storage
func (s Storage) Remove(p *storage.Page) error {
	fileName, err := fileName(p)
	if err != nil {
		return e.Wrap("can't remove file", err)
	}

	path := filepath.Join(s.basePath, p.UserName, fileName)

	if err := os.Remove(path); err != nil {
		return e.Wrap(fmt.Sprintf("can't remove file %s", path), err)
	}

	return nil
}

// IsExists checks if the specified page file exists in the storage
func (s Storage) IsExists(p *storage.Page) (bool, error) {
	fileName, err := fileName(p)
	if err != nil {
		return false, e.Wrap("can't check if file %s exists", err)
	}

	path := filepath.Join(s.basePath, p.UserName, fileName)

	switch _, err = os.Stat(path); {
	case errors.Is(err, os.ErrNotExist):
		return false, nil
	case err != nil:
		return false, e.Wrap(fmt.Sprintf("can't check if file %s exists", path), err)
	}

	return true, nil
}

// decodePage decodes and returns a Page object from the specified file
func (s Storage) decodePage(filePath string) (*storage.Page, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, e.Wrap("can't decode page", err)
	}

	defer func() { _ = f.Close() }()
	var p storage.Page

	if err := gob.NewDecoder(f).Decode(&p); err != nil {
		return nil, e.Wrap("can't decode page", err)
	}

	return &p, nil
}

// fileName generates a unique filename for the specified page based on its hash
func fileName(p *storage.Page) (string, error) {
	return p.Hash()
}
