package Solution

import (
	"fmt"
	"testing"
)

// TestSolution Example for solution test cases
func TestSolution(t *testing.T) {
	Push(1)
	Push(2)
	fmt.Println(Pop())
	fmt.Println(Pop())
}
