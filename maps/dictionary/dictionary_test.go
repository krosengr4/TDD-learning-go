package main

import "testing"

func TestSearch(t *testing.T) {

	d := Dictionary{"test": "this is just a test"}

	t.Run("known word", func(t *testing.T) {
		act, err := d.Search("test")

		assertDefinition(t, d, act, "test")
		assertNoError(t, err)
	})

	t.Run("unkown word", func(t *testing.T) {
		_, err := d.Search("unkown")

		assertError(t, err, ErrWordNotFound)
	})
}

func TestAdd(t *testing.T) {
	d := Dictionary{"test": "this is a test"}

	t.Run("adding new word", func(t *testing.T) {
		word := "new"
		definition := "this is a new word"
		err := d.Add(word, definition)

		assertError(t, err, nil)
		assertDefinition(t, d, definition, word)
	})

	t.Run("adding existing word", func(t *testing.T) {
		word := "test"
		def := "Koby Teeth :D"
		err := d.Add(word, def)

		assertError(t, err, ErrWordExists)
	})
}

func TestUpdate(t *testing.T) {
	d := Dictionary{"test": "this is a test"}

	t.Run("Updating existing word", func(t *testing.T) {
		word := "test"
		updatedDef := "this is an update test"
		err := d.Update(word, updatedDef)

		assertError(t, err, nil)
		assertDefinition(t, d, updatedDef, word)

	})

	t.Run("updating non existent word", func(t *testing.T) {
		word := "update"
		updatedDef := "gone til november, Wycleaf Jean"
		err := d.Update(word, updatedDef)

		assertError(t, err, ErrWordDoesNotExist)
	})
}

func TestDelete(t *testing.T) {
	d := Dictionary{"test": "test test test"}

	t.Run("deleting known word", func(t *testing.T) {
		word := "test"
		err := d.Delete(word)

		assertError(t, err, nil)

		_, err = d.Search(word)
		assertError(t, err, ErrWordNotFound)
	})

	t.Run("deleting unkown word", func(t *testing.T) {
		word := "dunno"
		err := d.Delete(word)

		assertError(t, err, ErrWordDoesNotExist)
	})
}

// Helper func to assert the dictionary definition
func assertDefinition(t testing.TB, d Dictionary, exp, word string) {
	t.Helper()

	act := d[word]

	if act != exp {
		t.Errorf("actual: %q expected: %q", act, exp)
	}
}

// Helper func to assert two strings are equal
func assertStrings(t testing.TB, act, exp string) {
	t.Helper()

	if act != exp {
		t.Errorf("actual: %q expected: %q", act, exp)
	}
}

// Helper func for asserting no errors
func assertNoError(t testing.TB, act error) {
	t.Helper()

	if act != nil {
		t.Fatal("got error when not expecting one")
	}
}

// Helper func to assert that an error is returned
func assertError(t testing.TB, act, exp error) {
	t.Helper()

	if act != exp {
		t.Errorf("actual: %q expected: %q", act, exp)
	}
}

