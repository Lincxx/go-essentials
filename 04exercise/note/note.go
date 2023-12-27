package note

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

// recall lowercase is private and upper case is public
// struct tags - all though we added tags this doesn't mean it will get picked up. We need a package to
// recognize it, like the Marshal package.
type Note struct {
	Title    string    `json:"title"`
	Content  string    `json:"content"`
	CreateAt time.Time `json:"created_at"`
}

func (note Note) Display() {
	fmt.Printf("Your note titled %v has the following content:\n\n%v\n\n", note.Title, note.Content)
}

// write to a file
func (note Note) Save() error {
	fileName := strings.ReplaceAll(note.Title, " ", "_")
	fileName = strings.ToLower(fileName) + ".json"

	// will convert data into json  - Marshal will convert and construct if publicly available (The Struct we are using)
	json, err := json.Marshal(note)
	if err != nil {
		return err
	}

	return os.WriteFile(fileName, json, 0644)
}

// simplier not to use a pointer on simple data structures
func New(title, content string) (Note, error) {
	if title == "" || content == "" {
		return Note{}, errors.New("invalid input")
	}

	return Note{
		Title:    title,
		Content:  content,
		CreateAt: time.Now(),
	}, nil
}
