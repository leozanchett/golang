package greetings

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

// Hello returns a greeting for the named person.
func Hello(name string) (string, error) {

	if name == "" {
		return "", errors.New("empty name")
	}

	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

// init is called before main.
func init() {
	fmt.Println(time.Now().UnixNano())
	rand.Seed(time.Now().UnixNano())
}

// randomFormat returns a random greeting format.
func randomFormat() string {
	formats := []string{
		"Greet to see you, %s!",
		"Hi, %s. Welcome!",
		"Hey, %s! Well met!",
		"Howdy, %s",
	}

	return formats[rand.Intn(len(formats))]
}

// Hellos returns a map that associates each of the named people
// with a greeting.
func Hellos(names []string) (map[string]string, error) {
	// A map to associate names with messages
	messages := make(map[string]string)
	// Loop through the received slice of names, calling
	// the Hello function to get a message for each name.
	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}
		// In the map, associate the name with the message.
		messages[name] = message
	}
	return messages, nil
}
