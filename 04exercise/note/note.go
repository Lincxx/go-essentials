package note

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

type Note struct {
	title    string
	content  string
	createAt time.Time
}

func (note Note) Display() {
	fmt.Printf("Your note titled %v has the following content:\n\n%v\n\n", note.title, note.content)
}

// write to a file
func (note Note) Save() {
	fileName := strings.ReplaceAll(note.title, " ", "_")
	fileName = strings.ToLower(fileName)
	os.WriteFile(fileName)
}

// simplier not to use a pointer on simple data structures
func New(title, content string) (Note, error) {
	if title == "" || content == "" {
		return Note{}, errors.New("invalid input")
	}

	return Note{
		title:    title,
		content:  content,
		createAt: time.Now(),
	}, nil
}
