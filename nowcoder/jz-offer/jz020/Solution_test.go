package Solution

import (
	"fmt"
	"testing"
)

// TestSolution Example for solution test cases
func TestSolution(t *testing.T) {
	s := MinSlack{}

	s.push(2)
	s.push(3)
	s.push(4)
	fmt.Println(s.min())
	s.pop()
	fmt.Println(s.min())
	s.push(1)
	fmt.Println(s.min())

}
