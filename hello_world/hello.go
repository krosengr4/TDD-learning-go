package main

import "fmt"

func Hello() string {
	return "Hello, world!"
}

func main() {
	hello := Hello()
	fmt.Println(hello)
}
