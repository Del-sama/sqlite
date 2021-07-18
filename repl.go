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

const (
	StatementInsert = iota
	StatementSelect
	StatementDelete
	StatementUpdate
)

const (
	PrepareStatementSuccess = iota
	PrepareStatementUnrecognized
)

type Statement struct {
	t int
}

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

// prepareStatement assigns recognized statements to the respective statement types
func prepareStatement(input string) (*Statement, int) {
	s := Statement{}
	switch {
	case strings.HasPrefix(input, "insert "):
		s.t = StatementInsert
		return &s, PrepareStatementSuccess
	case strings.HasPrefix(input, "select "):
		s.t = StatementSelect
		return &s, PrepareStatementSuccess
	case strings.HasPrefix(input, "update "):
		s.t = StatementUpdate
		return &s, PrepareStatementSuccess
	case strings.HasPrefix(input, "delete "):
		s.t = StatementDelete
		return &s, PrepareStatementSuccess
	default:
		fmt.Printf("Unrecognized keyword at start of '%s' ", input)
		return &s, PrepareStatementUnrecognized
	}
}

func executeStatement(statement *Statement) {
	switch statement.t {
	case StatementInsert:
		fmt.Println("Executing insert statement")
	case StatementSelect:
		fmt.Println("Executing select statement")
	case StatementUpdate:
		fmt.Println("Executing update statement")
	case StatementDelete:
		fmt.Println("Executing delete statement")
	}
}
