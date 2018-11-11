package Solution

import (
	"reflect"
	"testing"
)

func TestSolution(t *testing.T) {
	nodeG := TreeNode{Val: 1, Left: nil, Right: nil}
	nodeF := TreeNode{Val: 2, Left: &nodeG, Right: nil}
	nodeE := TreeNode{Val: 3, Left: nil, Right: nil}
	nodeD := TreeNode{Val: 4, Left: &nodeE, Right: nil}
	nodeC := TreeNode{Val: 5, Left: nil, Right: nil}
	nodeB := TreeNode{Val: 6, Left: &nodeD, Right: &nodeF}
	nodeA := TreeNode{Val: 7, Left: &nodeB, Right: &nodeC}

	result := preorderTraversal(&nodeA)
	result2 := preorderTraversal2(&nodeA)
	expected := []int{7, 6, 4, 3, 2, 1, 5}
	if !reflect.DeepEqual(result, expected) && !reflect.DeepEqual(result2, expected) {
		t.Fatalf("expected: %v, but got: %v",
			expected, result)
	}
}
