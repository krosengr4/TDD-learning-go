package main

import (
	"testing"
	"bytes"
)

// In this test we will be testing the actual printing using dependency injection

func TestGreet(t *testing.T) {

	// implements the writer interface
	buffer := bytes.Buffer{}
	Greet(&buffer, "Kevin")

	act := buffer.String()
	exp := "Hello, Kevin"

	if act != exp {
		t.Errorf("actual: %q expected: %q", act, exp)
	}
}
