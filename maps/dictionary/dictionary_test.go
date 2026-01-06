package main

import "testing"

func TestSearch(t *testing.T) {
	d := Dictionary{"test": "this is just a test"}

	act, err := d.Search("test")

	assertDefinition(t, d, act, "test")
	assertNoError(t, err)
}

// Helper func to assert the dictionary definition
func assertDefinition(t testing.TB, d Dictionary, exp, word string) {
	t.Helper()

	act := d[word]

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
