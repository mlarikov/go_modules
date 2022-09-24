package main

import (
	"fmt"

	"github.com/mlarikov/go_modules/greetings"
)

func main() {
	// Get a greeting message and print it.
	message := greetings.Hello("Mixa")
	fmt.Println(message)
}
