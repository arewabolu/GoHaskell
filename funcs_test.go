package gohaskell

import (
	"testing"
)

func TestPut(t *testing.T) {
	name := []string{"a", "b", "d", "e", "f", "g", "h", "i"}
	newName := Put(name, "c", 2)
	if newName[2] != "c" {
		t.Error("Could not put in string")
	}
}
