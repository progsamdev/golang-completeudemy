package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type Todo struct {
	Text string `json:"body"`
}

func New(text string) (*Todo, error) {

	if text == "" {
		return nil, errors.New("invalid input")
	}

	return &Todo{
		Text: text,
	}, nil
}

func (u *Todo) Display() {

	fmt.Printf("Your note titled has the following content: \n\n%v\n", u.Text)
}

func (u *Todo) Save() error {
	fileName := "todo.json"
	json, err := json.Marshal(u)

	if err != nil {
		return err
	}

	return os.WriteFile(fileName, json, 0644)
}
