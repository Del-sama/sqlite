package main

import (
	"os"
)

func main() {
	var table T = &Table{}
	t := table.createNewTable()
	repl(os.Stdin, t)
}
