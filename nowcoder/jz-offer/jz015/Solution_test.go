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
		expect *ListNode
	}{
		{"TestCase",
			&ListNode{Val: 1,
				Next: &ListNode{Val: 2,
					Next: &ListNode{Val: 3,
						Next: &ListNode{Val: 4,
							Next: &ListNode{Val: 5},
						},
					},
				},
			},
			&ListNode{Val: 5,
				Next: &ListNode{Val: 4,
					Next: &ListNode{Val: 3,
						Next: &ListNode{Val: 2,
							Next: &ListNode{Val: 1},
						},
					},
				},
			}},
	}

	//	Start test
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			actual := ReverseList(c.inputs)
			ast.Equal(c.expect, actual, "expected: %v, but actual: %v, with inputs: %v",
				c.expect, actual, c.inputs)
		})
	}
}
