package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/note/note"
	"example.com/note/todo"
)

type saver interface {
	Save() error
}

type outputtable interface {
	saver
	Display()
}

func main() {

	title, content := getNoteData()

	userNote, err := note.New(title, content)

	if err != nil {
		fmt.Print(err)
		return
	}

	todoText := getUserInput("Todo text: ")
	todo, err := todo.New(todoText)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = outputData(userNote)
	if err != nil {
		return
	}

	err = outputData(todo)

	if err != nil {
		return
	}

}

func getNoteData() (string, string) {
	title := getUserInput("note title: ")
	content := getUserInput("Note content: ")
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

func saveData(data saver) error {
	err := data.Save()
	if err != nil {
		fmt.Print("Saving the data failed.")
		return err
	}

	fmt.Print("Saving succeeded!")
	return nil
}

func outputData(data outputtable) error {
	data.Display()
	return saveData(data)
}

func printSomething(value any) { // value interface{}

	switch value.(type) {
	case int:
		fmt.Println("Integer: ", value)
	case float64:
		fmt.Println("Float: ", value)
	case string:
		fmt.Println("String: ", value)
	default:
		//do nothing
	}

}
