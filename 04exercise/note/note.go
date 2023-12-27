package note

import (
	"errors"
	"fmt"
	"time"
)

type Note struct {
	title    string
	content  string
	createAt time.Time
}

func (note Note) Display() {
	fmt.Printf("Your note titled %v has the following content:\n\n%v", note.title, note.content)
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
