package Solution

import (
	"testing"
)

// TestSolution Example for solution test cases
func TestSolution(t *testing.T) {
	obj := Constructor()
	obj.AppendTail(3)
	obj.DeleteHead()
	obj.DeleteHead()
}
