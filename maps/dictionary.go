package main

import "errors"

type Dictionary map[string]string

func (d Dictionary) Search(word string) (string, error) {
	if definition, ok := d[word]; ok {
		return definition, nil
	}
	return "", errors.New("could not find the word you were looking for")
}

func Search(dictionary map[string]string, word string) string {
	return dictionary[word]
}
