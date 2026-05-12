package mapCh

import "errors"

var (
	ErrorNotFound       = errors.New("not found")
	ErrWordExists       = errors.New("already exists")
	ErrWordDoesNotExist = errors.New("word does not exist yet")
)

type Dictionary map[string]string

func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]
	if !ok {
		return "", ErrorNotFound
	}
	return definition, nil
}

func (d Dictionary) Add(word string, definition string) error {
	_, err := d.Search(word)
	switch err {
	case ErrorNotFound:
		d[word] = definition
	case nil:
		return ErrWordExists
	default:
		return err
	}
	return nil
}

func (d Dictionary) Update(word string, definition string) error {
	_, err := d.Search(word)
	switch err {
	case ErrorNotFound:
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
	case ErrorNotFound:
		return ErrWordDoesNotExist
	case nil:
		delete(d, word)
	default:
		return err
	}
	return nil
}
