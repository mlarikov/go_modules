package greetings

import "fmt"

// Hello returns a greeting for the name in a message.
func Hello(name string) string {
	message := fmt.Sprintf("Hi %v. Welcome!", name)
	return message
}
