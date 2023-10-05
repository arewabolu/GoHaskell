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

func TestSum(t *testing.T) {
	name := []int{1, 2, 3}
	sum := Sum(name)
	t.Error(sum)
}

func TestContains(t *testing.T) {
	name := []int{1, 2, 3, 4}
	sum := Contains(name, 26)
	t.Error(sum)
}

func TestTake(t *testing.T) {
	name := []string{"a", "b", "d", "e", "f", "g", "h", "i"}
	nwName := Take(name, 26)
	t.Error(nwName)
	emptySlices := Take(name, 0)
	t.Error(emptySlices)
}

func TestDrop(t *testing.T) {
	name := []string{"a", "b", "d", "e", "f", "g", "h", "i"}
	nwName := Drop(name, 4)
	t.Error(nwName)
	trial2 := Drop(name, 0)
	t.Error(trial2)
	trial3 := Drop(name, 100)
	t.Error(trial3)
}
