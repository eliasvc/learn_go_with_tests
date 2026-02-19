package main

type Dictionary map[string]string

const (
	ErrNotFound         = DictionaryErr("could not find the word you were looking for")
	ErrWordExists       = DictionaryErr("cannot add word because it already exists")
	ErrWordDoesNotExist = DictionaryErr("cannot perform operation on word because it doesn't exist")
)

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

func (d Dictionary) Search(word string) (string, error) {
	if definition, ok := d[word]; ok {
		return definition, nil
	}
	return "", ErrNotFound
}

func (d Dictionary) Add(word, definition string) error {
	if _, ok := d[word]; ok {
		return ErrWordExists
	}

	d[word] = definition
	return nil
}

func (d Dictionary) Update(word, newDefinition string) error {
	if _, ok := d[word]; ok {
		d[word] = newDefinition
		return nil
	}

	return ErrWordDoesNotExist
}
