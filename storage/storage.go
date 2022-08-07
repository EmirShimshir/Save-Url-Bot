package storage

import (
	"Save_Url_Bot/lib/e"
	"crypto/sha1"
	"errors"
	"fmt"
	"io"
)

type Storage interface {
	Save(p *Page) error
	PickRandom(userName string) (*Page, error)
	Remove(p *Page) error
	IsExists(p *Page) (bool, error)
}

type Page struct {
	URL      string
	UserName string
}

var (
	ErrNoSavedPages = errors.New("no saved page")
)

func (p *Page) Hash() (string, error) {
	h := sha1.New()

	if _, err := io.WriteString(h, p.URL); err != nil {
		return "", e.Wrap("can't calculate hash", err)
	}

	if _, err := io.WriteString(h, p.UserName); err != nil {
		return "", e.Wrap("can't calculate hash", err)
	}

	// to convert a sha1 value to string, it needs to be encoded first,
	// since a hash is binary. The traditional encoding for SHA hashes is hex
	return fmt.Sprintf("%x", h.Sum(nil)), nil
}
