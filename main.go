package main

import (
	"os"
)

func main() {
	var table T = Table{}
	table.createNewTable()
	repl(os.Stdin)
}
