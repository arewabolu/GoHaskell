package gohaskell

import (
	"strings"
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
	nwName := Take(name, 4)
	t.Error(nwName)
	emptySlices := Take(name, 0)
	t.Error(emptySlices)
	nwslice := Take(name, 1)
	t.Error(nwslice)
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

func TestExclude(t *testing.T) {
	name := []string{"a", "b", "d", "e", "f", "g", "h", "i"}
	nwName := ExlcudeLast(name)
	t.Error(name, nwName)
}

func TestSpliter(t *testing.T) {
	word := "nine65fiveeightjmeight"
	texts := strings.ReplaceAll(word, "eight", "8")
	t.Error(texts)

	t.Run("slices", func(t *testing.T) {
		sep := strings.Split("De Minaur A.", " ")
		t.Error(sep, len(sep))
	})
}

func TestReverseTake(t *testing.T) {
	name := []string{"a", "b", "d", "e", "f", "g", "h", "i"}
	t.Run("testLargeLength", func(t *testing.T) {
		nwName := ReverseTake(name, 10)
		t.Error(nwName)
	})
	t.Run("TestOkLength", func(t *testing.T) {
		nwName := ReverseTake(name, 2)
		t.Error(nwName)
	})
}
