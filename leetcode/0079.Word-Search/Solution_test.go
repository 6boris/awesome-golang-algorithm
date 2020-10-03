package Solution

import (
	"testing"
)

func TestSolution(t *testing.T) {
	borad := [][]byte{
		[]byte{'A', 'B', 'C', 'E'},
		[]byte{'S', 'F', 'C', 'S'},
		[]byte{'A', 'D', 'E', 'E'},
	}
	// Given word = "ABCCED", return true.
	// Given word = "SEE", return true.
	// Given word = "ABCB", return false.
	t.Log("ABCCED=", exist(borad, "ABCCED"))
	t.Log("SEE=", exist(borad, "SEE"))
	t.Log("ABCB=", exist(borad, "ABCB"))
}
