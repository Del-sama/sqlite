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
)

// repl provides a repl for the db users to input statements
func repl(stdin io.Reader, table *T) {
	for {
		printPrompt()
		s := readInput(stdin)
		if strings.HasPrefix(s, ".") {
			handleMetaCommand(s)
		}
		stmnt, m := prepareStatement(s)
		if m == PrepareStatementSuccess {
			executeStatement(stmnt, table)
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

// handleMetaCommand executes meta commands
func handleMetaCommand(input string) {
	switch input {
	case ExitCommand:
		os.Exit(ExitSuccess)
	default:
		fmt.Printf("Unrecognized command %s \n ", input)
	}
}
