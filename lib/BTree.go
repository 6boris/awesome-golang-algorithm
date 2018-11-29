package Solution

import "fmt"

type BTree struct {
	root *TreeNode
}

type BTreeNode struct {
	data  interface{}
	left  *BTreeNode
	right *BTreeNode
}

func (t *BTree) MashBTree(s string) {
	fmt.Println("demo")
}
