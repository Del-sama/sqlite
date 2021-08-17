package main

import (
	"log"
)

func main() {
	var table T = &Table{}
	filename := "file.txt"

	_, _, err := table.dbOpen(filename)

	if err != nil {
		log.Fatalln(err)
	}
	// _ := os.Args[1:]

	// err = repl(scanner, os.Stdin, t)
	// if err != nil {
	// 	fmt.Println(err)
	// }
}
