package main

import (
	"fmt"
	"strings"
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

type Table struct {
	rowCount int
	rows     []Row
}

type T interface {
	createNewTable() *Table
	insertToTable(s *Statement) int
	selectAll() int
}

// prepareStatement assigns recognized statements to the respective statement types
func prepareStatement(input string) (*Statement, int) {
	s := Statement{}

	switch {
	case strings.HasPrefix(input, "insert "):
		return prepareInsert(s, input)
	case strings.HasPrefix(input, "select "):
		return prepareSelect(s)
	case strings.HasPrefix(input, "update "):
		return prepareUpdate(s)
	case strings.HasPrefix(input, "delete "):
		return prepareDelete(s)
	default:
		fmt.Printf("Unrecognized keyword at start of '%s' ", input)
		return &s, PrepareStatementUnrecognized
	}
}

func executeStatement(statement *Statement) int {
	var table T
	table = Table{}
	switch statement.StatementType {
	case StatementInsert:
		return table.insertToTable(statement)
	case StatementSelect:
		return table.selectAll()
	//case StatementUpdate:
	//	fmt.Println("Executing update statement")
	//case StatementDelete:
	//	fmt.Println("Executing delete statement")
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

func prepareUpdate(stmnt Statement) (*Statement, int) {
	stmnt.StatementType = StatementUpdate
	return &stmnt, PrepareStatementSuccess
}

func prepareDelete(stmnt Statement) (*Statement, int) {
	stmnt.StatementType = StatementDelete
	return &stmnt, PrepareStatementSuccess
}

func (t Table) insertToTable(s *Statement) int {
	t.rows = append(t.rows, s.InsertRow)
	t.rowCount += 1
	return ExecuteSuccess
}

func (t Table) selectAll() int {
	for _, row := range t.rows {
		fmt.Printf("%v \n", row)
	}
	return ExecuteSuccess
}

func (t Table) createNewTable() *Table {
	t.rowCount = 0
	t.rows = make([]Row, 0)
	return &t
}
