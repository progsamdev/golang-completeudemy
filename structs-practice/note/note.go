package note

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

type Note struct {
	Title     string    `json:"title"`
	Content   string    `json:"body"`
	CreatedAt time.Time `json:"created_at"`
}

func New(title, content string) (*Note, error) {

	if title == "" || content == "" {
		return nil, errors.New("invalid input")
	}

	return &Note{
		Title:     title,
		Content:   content,
		CreatedAt: time.Now(),
	}, nil
}

func (u *Note) Display() {

	fmt.Printf("Your note titled %v has the following content: \n\n%v\n", u.Title, u.Content)
}

func (u *Note) Save() error {
	fileName := strings.ReplaceAll(u.Title, " ", "_")
	fileName = strings.ToLower(fileName) + ".json"
	json, err := json.Marshal(u)

	if err != nil {
		return err
	}

	return os.WriteFile(fileName, json, 0644)
}
