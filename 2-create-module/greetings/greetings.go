package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

// Hello returns a greeting for the named person.
func HelloComplex(firstname string, lastname string) (string, error) {

	//can also make custom error type
	//also array of errors works

	if firstname == "" && lastname == "" {
		return "", errors.New("empty both names")
	}
	if firstname == "" {
		return "", errors.New("empty firstname")
	}
	if lastname == "" {
		return "", errors.New("empty lastname")
	}
	// Return a greeting that embeds the name in a message.
	message := fmt.Sprintf(randomFormat(), firstname, lastname)
	return message, nil
}
func Hello(name string) (string, error) {

	//can also make custom error type
	//also array of errors works

	if name == "" {
		return "", errors.New("empty name")
	}
	// Return a greeting that embeds the name in a message.
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

type ComplexNameType struct {
	fn string
	ln string
}

func CreateComplexName(fn string, ln string) ComplexNameType {
	var newName ComplexNameType
	newName.fn = fn
	newName.ln = ln
	return newName
}

func HelloMultipleComplex(nameVec []ComplexNameType) (map[ComplexNameType]string, error) {

	messages := make(map[ComplexNameType]string)

	for _, value := range nameVec {
		message, err := HelloComplex(value.fn, value.ln)
		if err != nil {
			return nil, err
		}
		messages[value] = message
	}
	return messages, nil
}

func HelloMultiple(nameVec []string) (map[string]string, error) {

	messages := make(map[string]string)

	for _, value := range nameVec {
		message, err := Hello(value)
		if err != nil {
			return nil, err
		}
		messages[value] = message
	}
	return messages, nil
}

func randomFormat() string {
	formats := []string{
		"Hi, %v %v. Welcome!",
		"Great to see you, %v %v!",
		"Hail, %v %v! Well met!",
	}
	return formats[rand.Intn(len(formats))]
}
