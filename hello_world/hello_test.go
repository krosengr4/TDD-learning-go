package main

import "testing"

// Test function must begin with the "Test" and take t *Testing as an arg
func TestHello(t *testing.T) {
	actual := Hello()
	expected := "Hello, world!"

	if actual != expected {
		t.Errorf("got %q but wanted %q", actual, expected)
	}
}
