package main

import (
	"os"
	"fmt"
)

func main() {
	var table T = Table{}
	t := table.createNewTable()
	fmt.Println(t)
	repl(os.Stdin, t)
}
