package main

import (
	"log"
	"os"
)

func main() {
	var table T = &Table{}
	filename := "file.txt"
	t, err := table.dbOpen(filename)
	if err != nil {
		log.Fatalln(err)
	}
	// _ := os.Args[1:]
	repl(os.Stdin, t)
}
