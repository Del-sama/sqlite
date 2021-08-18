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
func repl(stdin io.Reader, table *Table) error {
	for {
		printPrompt()
		s := readInput(stdin)
		if strings.HasPrefix(s, ".") {
			if err := handleMetaCommand(s, table); err != nil {
				return err
			}
		}
		stmnt, err := prepareStatement(s)
		if err != nil {
			return err
		}
		if err := executeStatement(stmnt, table); err != nil {
			return err
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
func handleMetaCommand(input string, table *Table) error {
	if input == ExitCommand {
		table.dbClose()
		os.Exit(ExitSuccess)
		return nil
	} else {
		return fmt.Errorf("unrecognized command %s \n ", input)
	}
}
