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
	definiton := d[word]
	return definiton, nil
}

/*
- Add func
- Search func
- Update func
*/
