package dictionary

import (
	"errors"
)

// alias
type Dictionary map[string]string

var (
	errorNotFound          = errors.New("Not found")
	errorWordAlreadyExists = errors.New("This word already exists")
	errorUpdateFailed      = errors.New("Update failed")
)

func (d Dictionary) Search(word string) (string, error) {
	value, exists := d[word]
	if exists {
		return value, nil
	}
	return "", errorNotFound
}

func (d Dictionary) Add(word string, def string) error {
	_, error := d.Search(word)
	switch error {
	case errorNotFound:
		d[word] = def
	case nil:
		return errorWordAlreadyExists
	}
	return nil
}

func (d Dictionary) Update(word string, def string) error {
	_, error := d.Search(word)
	switch error {
	case nil:
		d[word] = def
	case errorNotFound:
		return errorUpdateFailed
	}
	return nil
}

func (d Dictionary) Delete(word string) {
	delete(d, word)
}
