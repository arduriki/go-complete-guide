package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/note_taking/note"
)

func main() {
	title, content := getNoteData()

	userNote, err := note.New(title, content)
	if err != nil {
		fmt.Println(err)
		return
	}

	userNote.Display()
	err = userNote.Save()
	if err != nil {
		fmt.Println("Saving the note failed.")
		return
	}

	fmt.Println("Saving the note succeed!")
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
