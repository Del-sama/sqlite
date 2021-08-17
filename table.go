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
	dbOpen(filename string) (*bufio.Scanner, *Table, error)
	dbClose(scanner *bufio.Scanner) error
	getRowCount(scanner *bufio.Scanner) (int, error)
	saveRowCount(scanner *bufio.Scanner) error
	insertToTable(s *Statement) error
	selectAll(scanner *bufio.Scanner)
}

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

func executeStatement(scanner *bufio.Scanner, statement *Statement, table *Table) error {
	switch statement.StatementType {
	case StatementInsert:
		return table.insertToTable(statement)
	case StatementSelect:
		table.selectAll(scanner)
		return nil
	default:
		return fmt.Errorf("unrecognized statement %d ", statement.StatementType)
	}
}

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

func prepareSelect(stmnt Statement) (*Statement, error) {
	stmnt.StatementType = StatementSelect
	return &stmnt, nil
}

func (t *Table) dbOpen(filename string) (*bufio.Scanner, *Table, error) {
	f, err := os.OpenFile(filename, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0660)
	if err != nil {
		return nil, nil, err
	}
	t.file = f
	scanner := bufio.NewScanner(f)
	count, err := t.getRowCount(scanner)
	if err != nil {
		return nil, nil, err
	}
	t.rowCount = count
	return scanner, t, nil
}

func (t *Table) getRowCount(scanner *bufio.Scanner) (int, error) {
	var count []byte
	f := t.file
	info, err := f.Stat()
	if err != nil {
		return 0, err
	}
	size := info.Size()
	if size == 0 {
		f.WriteString("0\n")
		return 0, nil
	}

	f.ReadAt(count, 0)
	// var line int
	// for scanner.Scan() {
	// 	if line == 0 {
	// 		count, err := strconv.Atoi(scanner.Text())
	// 		if err != nil {
	// 			return 0, err
	// 		}
	// 		return count, nil
	// 	}
	// }
	fmt.Println("=============>", count)
	return 0, nil
}

func (t *Table) saveRowCount(scanner *bufio.Scanner) error {
	f := t.file
	count := t.rowCount

	var line int
	for scanner.Scan() {
		if line == 0 {
			en := json.NewEncoder(f)
			if err := en.Encode(count); err != nil {
				fmt.Println(err)
				return err
			}
			_, err := f.WriteString(strconv.Itoa(count))
			if err != nil {
				return err
			}
			return nil
		}
	}
	return nil
}

func (t *Table) dbClose(scanner *bufio.Scanner) error {
	f := t.file
	if err := t.saveRowCount(scanner); err != nil {
		return err
	}
	f.Close()
	return nil
}

func (t *Table) insertToTable(s *Statement) error {
	f := t.file
	en := json.NewEncoder(f)
	if err := en.Encode(s.InsertRow); err != nil {
		return err
	}
	t.rowCount += 1
	return nil
}

func (t *Table) selectAll(scanner *bufio.Scanner) {
	// look up scanner. Probably moves through thr file so it gets to the end
	// and can't go back. Perhaps scanner isn't what I need
	//  At saveRowCount scanner sees the first line and writes the rowcount to it
	//  but scubsequently, scanner does not
	// Also probably why select does not work until I exit. Scanner probably moves back to the top
	reader := bufio.NewReader(t.file)
	for {
		b, _, err := reader.ReadLine()
		fmt.Printf("%b \n", b)
		if err != nil {
			fmt.Println(err)
			break
		}
	}
	// var line int
	// for scanner.Scan() {
	// 	if line != 0 {
	// 		fmt.Printf("%v \n", scanner.Text())
	// 	}
	// 	line++
	// }
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
