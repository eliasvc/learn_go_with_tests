package maps

const (
	ErrNotFound         = DictionaryError("could not find the word you were looking for")
	ErrWordExists       = DictionaryError("word already exists")
	ErrWordDoesNotExist = DictionaryError("cannot perform operation on word because it does not exist")
)

type DictionaryError string

func (e DictionaryError) Error() string {
	return string(e)
}

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

func (d Dictionary) Update(word, definition string) error {

	_, err := d.Search(word)
	switch err {
	case ErrNotFound:
		return ErrWordDoesNotExist
	case nil:
		d[word] = definition
	default:
		return err
	}

	return nil
}

func (d Dictionary) Delete(word string) error {
	_, err := d.Search(word)
	switch err {
	case ErrNotFound:
		return ErrWordDoesNotExist
	case nil:
		delete(d, word)
	default:
		return err
	}

	return nil
}
