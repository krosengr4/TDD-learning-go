package main

import (
	"bytes"
	"fmt"
	"testing"
)

// In this test we will be testing the actual printing using dependency injection

func TestGreet(t *testing.T) {

	// implements the writer interface
	buffer := bytes.Buffer{}
	Greet(&buffer, "Kevin")

	act := buffer.String()
	exp := "Hello, Kevin"

	fmt.Println(exp)

	if act != exp {
		t.Errorf("actual: %q expected: %q", act, exp)
	}
}
