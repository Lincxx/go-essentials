package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

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
	err = userNote.Save()

	if err != nil {
		fmt.Println("Saving the note failed")
		return
	}

	fmt.Println("Saving the note....")
}

// helper func
func getUserInput(prompt string) string {
	fmt.Printf("%v ", prompt)
	//var value string
	//all scan funcs are for single words
	//fmt.Scanln(&value)

	//more complex solution - to read longer text input
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')

	if err != nil {
		return ""
	}

	//remove line break
	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")

	return text
}
