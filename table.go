package main

import (
	"errors"
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
	var username string
	var email string

	_, err := fmt.Sscanf(input, "insert %d %s %s", &i.id, &username, &email)

	if err != nil {
		return &s, PrepareStatementSyntaxError
	}
	if err = i.SetUsername(username); err != nil {
		fmt.Println(err)
		return &s, PrepareStatementSyntaxError
	}
	if err = i.SetEmail(email); err != nil {
		fmt.Println(err)
		return &s, PrepareStatementSyntaxError
	}
	s.InsertRow = i

	return &s, PrepareStatementSuccess
}

func prepareSelect(stmnt Statement) (*Statement, int) {
	stmnt.StatementType = StatementSelect
	return &stmnt, PrepareStatementSuccess
}

func (t *Table) insertToTable(s *Statement) int {
	t.rows = append(t.rows, s.InsertRow)
	t.rowCount += 1
	return ExecuteSuccess
}

func (t *Table) selectAll() int {
	for _, row := range t.rows {
		fmt.Printf("%v \n", row)
	}
	return ExecuteSuccess
}

func (t *Table) createNewTable() *Table {
	t.rowCount = 0
	t.rows = make([]Row, 0)
	return t
}

func (r *Row) SetUsername(username string) error {
	if len(username) > 32 {
		return errors.New("username is too long, max size is 32")
	}
	r.username = username
	return nil
}

func (r *Row) SetEmail(email string) error {
	if len(email) > 255 {
		return errors.New("email is too long, max size is 255")
	}
	r.email = email
	return nil
}
