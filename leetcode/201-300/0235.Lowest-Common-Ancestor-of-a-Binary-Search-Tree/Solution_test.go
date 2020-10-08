package Solution

import (
	"testing"
)

// 测试样例 root = [6,2,8,0,4,7,9,null,null,3,5], p = 2, q = 8
func TestSolution(t *testing.T) {
	tNode := &TreeNode{}
	tNode.val = 6

	tNode1 := &TreeNode{}
	tNode1.val = 2

	tNode2 := &TreeNode{}
	tNode2.val = 8

	tNode3 := &TreeNode{}
	tNode3.val = 0

	tNode4 := &TreeNode{}
	tNode4.val = 4

	tNode5 := &TreeNode{}
	tNode5.val = 7

	tNode6 := &TreeNode{}
	tNode6.val = 9

	tNode9 := &TreeNode{}
	tNode9.val = 3

	tNode10 := &TreeNode{}
	tNode10.val = 5

	// 第一层
	tNode.left = tNode1
	tNode.right = tNode2

	// 第二层
	tNode1.left = tNode3
	tNode1.right = tNode4
	tNode2.left = tNode5
	tNode2.right = tNode6

	// 第三层
	tNode4.left = tNode9
	tNode4.right = tNode10

	luckyNode := Lowest(tNode, tNode1, tNode2)
	t.Logf("luckyNode.value = %d", luckyNode.val)
	if luckyNode != tNode {
		t.Fatalf("Unlucky, false.")
	}
}
