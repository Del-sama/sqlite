package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func repl(stdin io.Reader) error {
	for {
		printPrompt()
		s := readInput(stdin)
		validateInput(s)
		os.Exit(0)
	}
}
func printPrompt() {
	fmt.Printf("db >")
}

func readInput(stdin io.Reader) string {
	reader := bufio.NewReader(stdin)
	s, _ := reader.ReadString('\n')
	return strings.TrimSpace(s)
}

func validateInput(input string) {
	if len(input) < 1 {
		fmt.Println("Error reading input")
		os.Exit(1)
	}
	if input != ".exit" {
		fmt.Printf("Unrecognized command %s", input)
	}
}
