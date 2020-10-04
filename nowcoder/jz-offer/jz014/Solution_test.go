package Solution

import (
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

// TestSolution Example for solution test cases
func TestSolution(t *testing.T) {
	ast := assert.New(t)
	//	test cases
	cases := []struct {
		name   string
		inputs *ListNode
		k      int
		expect *ListNode
	}{
		{"TestCase",
			&ListNode{Val: 1,
				Next: &ListNode{Val: 2,
					Next: &ListNode{Val: 3,
						Next: &ListNode{Val: 4,
							Next: &ListNode{Val: 5}},
					},
				},
			},
			5,
			&ListNode{Val: 1}},
	}

	//	Start test
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			actual := FindKthToTail(c.inputs, c.k)
			ast.Equal(c.expect.Val, actual.Val, "expected: %v, but actual: %v, with inputs: %v",
				c.expect, actual, c.inputs)
		})
	}
}

// TestSolution Example for solution test cases
func TestSolution2(t *testing.T) {
	ast := assert.New(t)
	//	test cases
	cases := []struct {
		name   string
		inputs *ListNode
		k      int
		expect *ListNode
	}{
		{"TestCase",
			&ListNode{Val: 1,
				Next: &ListNode{Val: 2,
					Next: &ListNode{Val: 3,
						Next: &ListNode{Val: 4,
							Next: &ListNode{Val: 5}},
					},
				},
			},
			5,
			&ListNode{Val: 1}},
	}

	//	Start test
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			actual := FindKthToTail2(c.inputs, c.k)
			ast.Equal(c.expect.Val, actual.Val, "expected: %v, but actual: %v, with inputs: %v",
				c.expect, actual, c.inputs)
		})
	}
}
