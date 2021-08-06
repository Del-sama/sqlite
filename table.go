package main

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strconv"
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
	Id       int
	Username string
	Email    string
}

type Table struct {
	file     *os.File
	rowCount int
}

type T interface {
	dbOpen(filename string) (*Table, error)
	dbClose()
	getRowCount(filename string) int
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

	_, err := fmt.Sscanf(input, "insert %d %s %s", &i.Id, &username, &email)

	if err != nil {
		return &s, PrepareStatementSyntaxError
	}
	if err = i.SetUsername(username); err != nil {
		return &s, PrepareStatementSyntaxError
	}
	if err = i.SetEmail(email); err != nil {
		return &s, PrepareStatementSyntaxError
	}
	s.InsertRow = i

	return &s, PrepareStatementSuccess
}

func prepareSelect(stmnt Statement) (*Statement, int) {
	stmnt.StatementType = StatementSelect
	return &stmnt, PrepareStatementSuccess
}

func (t *Table) dbOpen(filename string) (*Table, error) {
	f, err := os.OpenFile(filename, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0660)
	if err != nil {
		return nil, err
	}
	t.file = f
	t.rowCount = t.getRowCount()
	return t, nil
}

func (t *Table) getRowCount() int {
	f := t.file
	scanner := bufio.NewScanner(f)
	var line int
	for scanner.Scan() {
		if line == 0 {
			count, err := strconv.Atoi(scanner.Text())
			if err != nil {
				fmt.Println(err)
			}
			return count
		}
	}
	return 0
}

func (t *Table) dbClose() {
	f := t.file
	f.Close()
}

func (t *Table) insertToTable(s *Statement) int {
	f := t.file
	en := json.NewEncoder(f)
	if err := en.Encode(s.InsertRow); err != nil {
		return ExecuteFailure
	}
	t.rowCount += 1
	return ExecuteSuccess
}

func (t *Table) selectAll() int {
	f := t.file
	scanner := bufio.NewScanner(f)
	var line int
	for scanner.Scan() {
		if line != 0 {
			fmt.Printf("%v \n", scanner.Text())
		}
		line++
	}
	return ExecuteSuccess
}

func (r *Row) SetUsername(username string) error {
	if len(username) > 32 {
		return errors.New("username is too long, max size is 32")
	}
	r.Username = username
	return nil
}

func (r *Row) SetEmail(email string) error {
	if len(email) > 255 {
		return errors.New("email is too long, max size is 255")
	}
	r.Email = email
	return nil
}
