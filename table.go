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

type (
	// Statement ...
	Statement struct {
		StatementType int
		InsertRow     Row
	}

	// Row ...
	Row struct {
		Id       int
		Username string
		Email    string
	}

	// Table ...
	Table struct {
		files    map[string]*os.File
		rowCount int
	}

	// DB ...
	DB interface {
		dbOpen(filename string) (*Table, error)
		dbClose() error
		getRowCount() (int, error)
		saveRowCount() error
		insertToTable(s *Statement) error
		selectAll()
	}
)

// prepareStatement assigns recognized statements to the respective statement types
func prepareStatement(input string) (*Statement, error) {
	s := Statement{}
	switch {
	case strings.HasPrefix(input, "insert"):
		return prepareInsert(s, input)
	case strings.HasPrefix(input, "select"):
		return prepareSelect(s)
	default:
		return &s, fmt.Errorf("unrecognized keyword at start of `%s`", input)
	}
}

// executeStatemnt ...
func executeStatement(statement *Statement, table *Table) error {
	switch statement.StatementType {
	case StatementInsert:
		return table.insertToTable(statement)
	case StatementSelect:
		table.selectAll()
		return nil
	default:
		return fmt.Errorf("unrecognized statement %d ", statement.StatementType)
	}
}

// prepareInsert prepares the insert statement
func prepareInsert(s Statement, input string) (*Statement, error) {
	s.StatementType = StatementInsert
	i := Row{}
	var username string
	var email string

	_, err := fmt.Sscanf(input, "insert %d %s %s", &i.Id, &username, &email)

	if err != nil {
		return nil, err
	}
	if err = i.SetUsername(username); err != nil {
		return nil, err
	}
	if err = i.SetEmail(email); err != nil {
		return nil, err
	}
	s.InsertRow = i

	return &s, nil
}

// prepareSelect prepares the select statement
func prepareSelect(stmnt Statement) (*Statement, error) {
	stmnt.StatementType = StatementSelect
	return &stmnt, nil
}

// dbOpen opens the db
func (t *Table) dbOpen(filename string) (*Table, error) {
	table, err := os.OpenFile(fmt.Sprintf("%s.txt", filename), os.O_RDWR|os.O_APPEND|os.O_CREATE, 0660)
	if err != nil {
		return nil, err
	}

	rowCount, err := os.OpenFile(fmt.Sprintf("%s_%s.txt", filename, "count"), os.O_RDWR|os.O_CREATE, 0660)
	if err != nil {
		return nil, err
	}

	t.files = map[string]*os.File{"table": table, "rowCount": rowCount}

	count, err := t.getRowCount()
	if err != nil {
		return nil, err
	}

	t.rowCount = count
	return t, nil
}

// dbClose shuts down the db
func (t *Table) dbClose() error {
	table := t.files["table"]
	count := t.files["rowCount"]
	if err := t.saveRowCount(); err != nil {
		return err
	}
	table.Close()
	count.Close()
	return nil
}

// insertToTable write rows to the table
func (t *Table) insertToTable(s *Statement) error {
	f := t.files["table"]
	en := json.NewEncoder(f)
	if err := en.Encode(s.InsertRow); err != nil {
		return err
	}

	t.rowCount += 1
	return nil
}

// selectAll returns all rows written to the table
func (t *Table) selectAll() {
	f := t.files["table"]
	var line int
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		fmt.Printf("%v \n", scanner.Text())
		line++
	}
}

// gerRowCount gets the row count from file
func (t *Table) getRowCount() (int, error) {
	var count int
	f := t.files["rowCount"]

	info, err := f.Stat()
	if err != nil {
		return 0, err
	}

	size := info.Size()
	if size == 0 {
		_, err = fmt.Fprintf(f, "%d", 0)
		if err != nil {
			return 0, err
		}
	}

	var line int
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		if line == 0 {
			count, err := strconv.Atoi(scanner.Text())
			if err != nil {
				return 0, err
			}
			return count, nil
		}
	}
	return count, nil
}

// saveRowCount saves the row count to file
func (t *Table) saveRowCount() error {
	f := t.files["rowCount"]
	count := t.rowCount

	err := f.Truncate(0)
	if err != nil {
		return err
	}
	_, err = f.Seek(0, 0)
	if err != nil {
		return err
	}
	_, err = fmt.Fprintf(f, "%d", count)
	if err != nil {
		return err
	}
	return nil
}

// SetUsername is a setter method for Username
func (r *Row) SetUsername(username string) error {
	if len(username) > 32 {
		return errors.New("username is too long, max size is 32")
	}
	r.Username = username
	return nil
}

// SetEmail is a setter method for Email
func (r *Row) SetEmail(email string) error {
	if len(email) > 255 {
		return errors.New("email is too long, max size is 255")
	}
	r.Email = email
	return nil
}
