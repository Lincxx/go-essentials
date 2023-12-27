package main

import (
	"fmt"

	"example.com/exercise/note"
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

	userNote.Display()
}

// helper func
func getUserInput(prompt string) string {
	fmt.Println(prompt)
	var value string
	//all scan methods are for single words
	//fmt.Scanln(&value)

	//more complex saolution

	return value
}
