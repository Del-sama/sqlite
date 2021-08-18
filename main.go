package main

import (
	"fmt"
	"log"
	"os"
	"unicode"
)

func main() {
	var table DB = &Table{}
	filename := os.Args[1]

	if !isValidFileName(filename) {
		log.Fatalln("Invalid filename: filename must not contain letters or numbers")
	}

	t, err := table.dbOpen(filename)

	if err != nil {
		log.Fatalln(err)
	}
	err = repl(os.Stdin, t)
	if err != nil {
		fmt.Println(err)
	}
}

func isValidFileName(fileName string) bool {
	// Make sure filename does not contain special characters or numbers
	if len(fileName) < 1 {
		return false
	}
	for _, val := range fileName {
		if !unicode.IsLetter(val) {
			return false
		}
	}
	return true
}
