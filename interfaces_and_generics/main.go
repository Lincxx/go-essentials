package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/note/note"
	"example.com/note/todo"
)

// interface - naming convention -  method name plus er
type saver interface {
	Save() error
}

func main() {
	title, content := getNoteData()
	todoText := getUserInput("Todo text")

	userNote, err := note.New(title, content)

	if err != nil {
		fmt.Println(err)
		return
	}

	todo, err := todo.New(todoText)

	if err != nil {
		fmt.Println(err)
		return
	}

	todo.Display()
	err = saveData(todo)

	if err != nil {
		fmt.Println("Saving the note failed.")
		return
	}

	userNote.Display()
	err = saveData(userNote)

	if err != nil {
		return
	}

	fmt.Println("Saving the todo succeeded!")
	fmt.Println("Saving the note succeeded!")
}

func saveData(data saver) error {
	//this would be todo.Save or note.Save or whatever you pass in
	// because our todo and note struct has a func called Save that doesn't take in any args and adheres to the interface
	err := data.Save()

	if err != nil {
		fmt.Println("Saving the note failed.")
		return err
	}

	return nil
}

func getNoteData() (string, string) {
	title := getUserInput("Note title:")
	content := getUserInput("Note content:")

	return title, content
}

func getUserInput(prompt string) string {
	fmt.Printf("%v ", prompt)

	reader := bufio.NewReader(os.Stdin)

	text, err := reader.ReadString('\n')

	if err != nil {
		return ""
	}

	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")

	return text
}
