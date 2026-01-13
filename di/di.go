package main

import (
	"fmt"
	"os"
	"io"
)

func Greet(writer io.Writer, name string) {
	// Fprint returns number of bytes written so that we can test what prints
	fmt.Fprintf(writer, "Hello, %s", name)
}

func main() {
	Greet(os.Stdout, "Kevin")
}
