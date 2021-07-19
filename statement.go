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

type Statement struct {
	t int
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
	case strings.HasPrefix(input, "insert "):
		return handleInsert(s, input)
	case strings.HasPrefix(input, "select "):
		return handleSelect(s, input)
	case strings.HasPrefix(input, "update "):
		return handleUpdate(s, input)
	case strings.HasPrefix(input, "delete "):
		return handleDelete(s, input)
	default:
		fmt.Printf("Unrecognized keyword at start of '%s' ", input)
		return &s, PrepareStatementUnrecognized
	}
}

func executeStatement(statement *Statement) {
	switch statement.t {
	case StatementInsert:
		fmt.Println("Executing insert statement")
	case StatementSelect:
		fmt.Println("Executing select statement")
	case StatementUpdate:
		fmt.Println("Executing update statement")
	case StatementDelete:
		fmt.Println("Executing delete statement")
	}
}

func handleInsert(stmnt Statement, input string) (*Statement, int) {
	stmnt.t = StatementInsert
	i := Row{}
	_, err := fmt.Sscanf(input, "insert %d %s %s", &i.id, &i.username, &i.email)
	if err != nil {
		return &stmnt, PrepareStatementSyntaxError
	}
	return &stmnt, PrepareStatementSuccess
}

func handleSelect(stmnt Statement, input string) (*Statement, int) {
	stmnt.t = StatementSelect
	return &stmnt, PrepareStatementSuccess
}

func handleUpdate(stmnt Statement, input string) (*Statement, int) {
	stmnt.t = StatementUpdate
	return &stmnt, PrepareStatementSuccess
}

func handleDelete(stmnt Statement, input string) (*Statement, int) {
	stmnt.t = StatementDelete
	return &stmnt, PrepareStatementSuccess
}

+const uint32_t ID_SIZE = size_of_attribute(Row, id);
+const uint32_t USERNAME_SIZE = size_of_attribute(Row, username);
+const uint32_t EMAIL_SIZE = size_of_attribute(Row, email);
+const uint32_t ID_OFFSET = 0;
+const uint32_t USERNAME_OFFSET = ID_OFFSET + ID_SIZE;
+const uint32_t EMAIL_OFFSET = USERNAME_OFFSET + USERNAME_SIZE;
+const uint32_t ROW_SIZE = ID_SIZE + USERNAME_SIZE + EMAIL_SIZE;
