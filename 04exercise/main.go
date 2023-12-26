package main

import (
	"example.com/exercise/note"
	"fmt"
)

func getNoteData() (string, string) {

	title := getUserInput("Note title:")
	content := getUserInput("Note Content:")

	return title, content
}

func main() {
	title, content := getNoteData()

	userNote, err := note.New(title, content)

	if err != nil {
		fmt.Println(err)
		return
	}
}

// helper func
func getUserInput(prompt string) string {
	fmt.Println(prompt)
	var value string
	fmt.Scanln(&value)

	return value
}
