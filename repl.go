package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

const (
	ExitCommand = ".exit"
	ExitSuccess = 0
	ExitFailure = 1
)

// repl provides a repl for the db users to input statements
func repl(stdin io.Reader) {
	for {

		printPrompt()
		s := readInput(stdin)
		if isValidInput(s) {
			if strings.HasPrefix(s, ".") {
				handleMetaCommand(s)
			}
			s, m := prepareStatement(s)
			if m == PrepareStatementSuccess {
				executeStatement(s)
			} else {
				os.Exit(ExitFailure)
			}

		}

	}

}

// printPrompt prints a cmd prompt to users
func printPrompt() {
	fmt.Printf("db > ")
}

// readInput reads std input and trimspaces around the input
func readInput(stdin io.Reader) string {
	reader := bufio.NewReader(stdin)
	s, _ := reader.ReadString('\n')
	return strings.TrimSpace(s)
}

// isValidInput checks that an input is valid
func isValidInput(input string) bool {
	if len(input) < 1 {
		fmt.Println("Error reading input")
		return false
	}
	return true
}

// handleMetacommand executes meta commands
func handleMetaCommand(input string) {
	switch input {
	case ExitCommand:
		os.Exit(ExitSuccess)
	default:
		fmt.Printf("Unrecognized command %s \n ", input)
	}
}
