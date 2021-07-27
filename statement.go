package main

import (
	"fmt"
	"strings"
)

const (
	StatementInsert = iota
	StatementSelect
)

const (
	PrepareStatementSuccess = iota
	PrepareStatementUnrecognized
	PrepareStatementSyntaxError
)

const (
	ExecuteSuccess = iota
	ExecuteFailure
)

type Statement struct {
	StatementType int
	InsertRow     Row
}

type Row struct {
	id       int
	username string
	email    string
}

// prepareStatement assigns recognized statements to the respective statement types
func prepareStatement(input string) (*Statement, int) {
	s := Statement{}
	switch {
	case strings.HasPrefix(input, "insert"):
		return prepareInsert(s, input)
	case strings.HasPrefix(input, "select"):
		return prepareSelect(s)
	default:
		fmt.Printf("Unrecognized keyword at start of '%s' \n", input)
		return &s, PrepareStatementUnrecognized
	}
}

func executeStatement(statement *Statement, table *Table) int {
	switch statement.StatementType {
	case StatementInsert:
		return table.insertToTable(statement)
	case StatementSelect:
		return table.selectAll()
	default:
		fmt.Println("Unrecognized statement")
		return ExecuteFailure
	}
}

func prepareInsert(s Statement, input string) (*Statement, int) {
	s.StatementType = StatementInsert
	i := Row{}
	_, err := fmt.Sscanf(input, "insert %d %s %s", &i.id, &i.username, &i.email)
	s.InsertRow = i
	fmt.Println(err)
	fmt.Println(s)

	if err != nil {
		return &s, PrepareStatementSyntaxError
	}
	return &s, PrepareStatementSuccess
}

func prepareSelect(stmnt Statement) (*Statement, int) {
	stmnt.StatementType = StatementSelect
	return &stmnt, PrepareStatementSuccess
}
