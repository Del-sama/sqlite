package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

// repl provides a repl for the db users to input statements
func repl(stdin io.Reader) {
	for {
		printPrompt()
		s := readInput(stdin)
		validateInput(s)
	}
}

// printPrompt prints a cmd prompt to users
func printPrompt() {
	fmt.Printf("db >")
}

// readInput reads std input and trimspaces around the input
func readInput(stdin io.Reader) string {
	reader := bufio.NewReader(stdin)
	s, _ := reader.ReadString('\n')
	return strings.TrimSpace(s)
}

// validateInput validates that input is validate and recognized
func validateInput(input string) {
	if len(input) < 1 {
		fmt.Println("Error reading input")
		os.Exit(1)
	}
	if input != ".exit" {
		fmt.Printf("Unrecognized command %s", input)
	}
	os.Exit(0)
}
