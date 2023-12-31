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

// type displayer interface {
// 	Display()
// }

// embedded interface
type outputtable interface {
	saver
	Display()
}

// type outputtable interface {
// 	Save() error
// 	Display()
// }

func main() {
	printSomething(1)
	printSomething("1")
	printSomething(1.2)

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

	printSomething(todo)

	// todo.Display()
	// err = saveData(todo)
	err = outputData(todo)

	if err != nil {
		fmt.Println("Saving the note failed.")
		return
	}

	// userNote.Display()
	// err = saveData(userNote)
	err = outputData(userNote)

	//we could get rid of the err check, since if there was an error the program would stop
	// if err != nil {
	// 	return
	// }

	fmt.Println("Saving the todo succeeded!")
	fmt.Println("Saving the note succeeded!")
}

// he Special "Any Value Allowed" Type
func printSomething(value interface{}) {
	//Extracting Type Information From Values
	intVal, ok := value.(int)
	if ok {
		fmt.Println("Integer: ", intVal)
		return
	}

	floatVal, ok := value.(float64)
	if ok {
		fmt.Println("Integer: ", floatVal)
		return
	}

	strVal, ok := value.(string)
	if ok {
		fmt.Println("Integer: ", strVal)
		return
	}

	// switch value.(type) {
	// case int:
	// 	fmt.Println("Int: ", value)
	// case float64:
	// 	fmt.Println("Float64: ", value)
	// case string:
	// 	fmt.Println("String: ", value)
	// }

}

func outputData(data outputtable) error {
	data.Display()
	return saveData(data)

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
