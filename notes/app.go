package main

import (
	"bufio"
	"com.example/notes/note"
	"fmt"
	"os"
	"strings"
)

func main() {

	title, content := getNoteData()

	userNote, err := note.New(title, content)

	if err != nil {
		fmt.Println(err)
		return
	}

	userNote.DisplayNote()

	err = userNote.Save()

	if err != nil {
		fmt.Println("Saving the note failed: ", err)
		return
	}

	fmt.Println("Note saved successfully!")

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

func getNoteData() (string, string) {

	title := getUserInput("Note Title: ")

	content := getUserInput("Note Content: ")

	return title, content

}
