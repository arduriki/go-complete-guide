package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/note_taking/note"
	"example.com/note_taking/todo"
)

// Creating interfaces: a contract that guarantees that a certain value has a certain method.
// Whichever struct "signs the contract" would have the following methods.
// Indicate the kind of values that the funcs would accept as inputs.
// "saver" due to only one method + 'er'.
type saver interface {
	Save() error
}

// Embed other interfaces.
type outputtable interface {
	// Include the 'saver' interface.
	saver
	Display()
}

func main() {
	// Example of "any value allowed"
	printSomething(1)
	printSomething(1.5)
	printSomething("Hello")

	title, content := getNoteData()
	todoText := getUserInput("Todo text: ")

	todo, err := todo.New(todoText)
	if err != nil {
		fmt.Println(err)
		return
	}

	printSomething(todo)

	userNote, err := note.New(title, content)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Display and save a todo.
	err = outputData(todo)
	if err != nil {
		return
	}

	// Display and save a note.
	outputData(userNote)
}

// Example of "any value allowed"
// func printSomething(value any) {
// 	fmt.Println(value)
// }

// Another kind of "any kind of value".
func printSomething(value interface{}) {
	// Type value and a bool
	intVal, ok := value.(int)
	if ok {
		fmt.Println("Integer:", intVal)
		return
	}

	floatVal, ok := value.(float64)
	if ok {
		fmt.Println("Float:", floatVal)
		return
	}

	stringVal, ok := value.(string)
	if ok {
		fmt.Println("String:", stringVal)
		return
	}
	
	// Another kind of 'switch' statement: checking the type of value.
	// switch value.(type) {
	// case int:
	// 	fmt.Println("Interger:", value)
	// case float64:
	// 	fmt.Println("Float:", value)
	// case string:
	// 	fmt.Println(value)
	 	// It isn't necessary to specify a fallback option with 'default'.
	// }
}

func outputData(data outputtable) error {
	data.Display()
	return saveData(data)
}

// Using the interface: guaranteeing that the data is the one indicated in the interface.
func saveData(data saver) error {
	err := data.Save()
	if err != nil {
		fmt.Println("Saving the note failed.")
		return err
	}

	fmt.Println("Saving the note succeed!")

	return nil
}

func getNoteData() (string, string) {
	title := getUserInput("Note title:")
	content := getUserInput("Note content:")

	return title, content
}

func getUserInput(prompt string) string {
	fmt.Printf("%v ", prompt)
	// To enter larger texts for user input.
	reader := bufio.NewReader(os.Stdin)

	// Stop reading at line break byte.
	text, err := reader.ReadString('\n')
	if err != nil {
		return ""
	}

	// Remove some text at the end.
	text = strings.TrimSuffix(text, "\n")
	// Remove Window's line break.
	text = strings.TrimSuffix(text, "\r")

	return text
}
