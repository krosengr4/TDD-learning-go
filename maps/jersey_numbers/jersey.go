package main

import (
	"errors"
)

// Error vars
var (
	ErrNumberNotFound     = errors.New("error: could not find a player with that jersey number")
	ErrNumbExists         = errors.New("error: a player with that jersey number already exists")
	ErrNumberDoesNotExist = errors.New("error: could not perform operation, player with that number does not exist")
)

// Map to hold the numbers and the players
type Roster map[int]string

func (r Roster) Search(number int) (int, string, error) {

	player, ok := r[number]
	if !ok {
		return 0, "", ErrNumberNotFound
	}

	return number, player, nil
}

func (r Roster) Add(number int, name string) error {

	_, _, err := r.Search(number)
	switch err {
	case ErrNumberNotFound:
		r[number] = name
	case nil:
		return ErrNumbExists
	default:
		return err
	}

	return nil
}

func (r Roster) Update(number int, newName string) (string, error) {
	_, _, err := r.Search(number)
	switch err {
	case ErrNumberNotFound:
		return "", ErrNumberDoesNotExist
	case nil:
		r[number] = newName
	default:
		return "", err
	}

	return r[number], nil
}

func (r Roster) Delete(number int) error {
	_, _, err := r.Search(number)
	switch err {
	case ErrNumberNotFound:
		return ErrNumberDoesNotExist
	case nil:
		delete(r, number)
	default:
		return err
	}

	return nil
}
