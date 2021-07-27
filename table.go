package main

import "fmt"

type Table struct {
	rowCount int
	rows     []Row
}

type T interface {
	createNewTable() *Table
	insertToTable(s *Statement) int
	selectAll() int
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
