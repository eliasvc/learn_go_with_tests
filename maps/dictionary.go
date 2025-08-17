package maps

import "errors"

var ErrNotFound = errors.New("could not find the word you were looking for")
var ErrWordExists = errors.New("word already exists")

type Dictionary map[string]string

func (d Dictionary) Search(key string) (string, error) {
	if val, ok := d[key]; ok {
		return val, nil
	}
	return "", ErrNotFound
}

func (d Dictionary) Add(word, definition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		d[word] = definition
	case nil:
		return ErrWordExists
	default:
		return err
	}

	return nil
}
