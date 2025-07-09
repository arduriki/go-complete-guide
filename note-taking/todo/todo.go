package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type Todo struct {
	Text string `json:"text"`
}

func (todo Todo) Display() {
	fmt.Println(todo.Text)
}

func (todo Todo) Save() error {
	// Write a file name.
	fileName := "todo.json"

	// To marshal we need the struct's values to be public
	json, err := json.Marshal(todo)
	if err != nil {
		return err
	}

	return os.WriteFile(fileName, json, 0644)
}

func New(content string) (Todo, error) {
	if content == "" {
		// return an empty note and the error
		return Todo{}, errors.New("[ERROR] invalid input")
	}

	// if all correct, return a note and empty error.
	return Todo{
		Text: content,
	}, nil
}
