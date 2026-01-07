package main

import "errors"

// Errors
var (
	// Definition could not be found for the given word (for Search func)
	ErrWordNotFound = errors.New("error: that word is not in the dictionary!")

	// Word already exists in the dictionary (for Add func)
	ErrWordExists = errors.New("error: cannot add word that already exists!")

	// Word does not exist in the dictionary (for Update func)
	ErrWordDoesNotExist = errors.New("error: cannot perform operation, word does not exist!")
)

// Map to store definitions to words
type Dictionary map[string]string

func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]
	if !ok {
		return "", ErrWordNotFound
	}

	return definition, nil
}

func (d Dictionary) Add(word, definition string) error {

	_, err := d.Search(word)
	switch err {
	case ErrWordNotFound:
		d[word] = definition
	case nil:
		return ErrWordExists
	default:
		return err
	}

	return nil
}

func (d Dictionary) Update(word, newDefinition string) error {

	_, err := d.Search(word)

	switch err {
	case ErrWordNotFound:
		return ErrWordDoesNotExist
	case nil:
		d[word] = newDefinition
	default:
		return err
	}

	return nil
}

func (d Dictionary) Delete(word string) error {

	_, err := d.Search(word)

	switch err {
	case ErrWordNotFound:
		return ErrWordDoesNotExist
	case nil:
		delete(d, word)
	default:
		return err
	}

	return nil
}

