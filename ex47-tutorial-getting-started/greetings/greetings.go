package greetings

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

// Go executes init functions automatically at program startup, after global variables have been initialized.
func init() {
	rand.Seed(time.Now().UnixNano())
}

// Hello returns a greeting for the named person.
func Hello(name string) (string, error) {
	formats := []string{
		"Hi, %v.",
		"Hello, %v.",
		"%v, how are you?",
	}
	if name == "" {
		return "", errors.New("empty name")
	}
	// Return a greeting that embeds the name in a message.
	message := fmt.Sprintf(formats[rand.Intn(len(formats))], name)
	return message, nil
}

func Hellos(names []string) (map[string]string, error) {
	r := make(map[string]string)
	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}
		r[name] = message
	}
	return r, nil
}
