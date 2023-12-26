package main

import (
	"errors"
	"fmt"
)

func getNoteData() (string, string, error) {
	title, err := getUserInput("Note title:")
	if err != nil {
		return "", "", err
	}

	content, err := getUserInput("Note Content")
	if err != nil {
		return "", "", err
	}

	return title, content, nil
}

func main() {
	title, content, err := getNoteData()

	if err != nil {
		fmt.Println(err)
	}

}

// helper func
func getUserInput(prompt string) (string, error) {
	fmt.Println(prompt)
	var value string
	fmt.Scanln(&value)

	if value == "" {
		return "", errors.New("invalid input")
	}

	return value, nil
}
