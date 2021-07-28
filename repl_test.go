package main

import (
	"bytes"
	"testing"
)

func TestRepl(t *testing.T) {
	var stdin bytes.Buffer

	stdin.Write([]byte("test\n"))
	res := readInput(&stdin)

	if res != "test" {
		t.Errorf("Expected %s, got %s", "test", res)
	}

	stdin.Write([]byte(".start\n"))
	res = readInput(&stdin)

	if res != "10" {
		t.Errorf("Expected %s, got %s", "10", res)
	}
}
