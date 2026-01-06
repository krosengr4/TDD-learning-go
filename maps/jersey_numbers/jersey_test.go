package main

import "testing"

func TestSearch(t *testing.T) {
	roster := Roster{1: "Test Test"}

	t.Run("known player", func(t *testing.T) {
		_, act, _ := roster.Search(1)
		exp := "Test Test"

		assertString(t, act, exp)
	})

	t.Run("unknown player", func(t *testing.T) {
		_, _, act := roster.Search(2)

		assertError(t, act, ErrNumberNotFound)
	})
}

func TestAdd(t *testing.T) {
	roster := Roster{1: "Test Test"}

	t.Run("Adding new player", func(t *testing.T) {
		number := 2
		name := "New Player"
		err := roster.Add(number, name)

		assertError(t, err, nil)
		assertName(t, roster, number, name)
	})

	t.Run("Adding existing player", func(t *testing.T) {
		number := 1
		name := "Existing Player"
		err := roster.Add(number, name)

		assertError(t, err, ErrNumbExists)
	})
}

func TestUpdate(t *testing.T) {
	roster := Roster{1: "Test Test"}

	t.Run("Updating existing player", func(t *testing.T) {
		number := 1
		newName := "Test Update"
		_, err := roster.Update(number, newName)

		assertError(t, err, nil)
		assertName(t, roster, number, newName)
	})

	t.Run("Update non existing player", func(t *testing.T) {
		number := 3
		newName := "Test Update"
		_, err := roster.Update(number, newName)

		assertError(t, err, ErrNumberDoesNotExist)
	})
}

// Helper func to assert actual and expected values
func assertString(t testing.TB, act, exp string) {
	t.Helper()

	if act != exp {
		t.Errorf("actual: %q expected: %q", act, exp)
	}
}

// Helper func to assert player name was added
func assertName(t testing.TB, roster Roster, number int, exp string) {
	t.Helper()

	_, act, err := roster.Search(number)
	if err != nil {
		t.Fatal("should have found player name and did not")
	}

	if act != exp {
		t.Errorf("actual: %q expected: %q", act, exp)
	}
}

// Helper func to assert errors
func assertError(t testing.TB, act, exp error) {
	t.Helper()

	if act != exp {
		t.Errorf("actual: %q expected: %q", act, exp)
	}
}
