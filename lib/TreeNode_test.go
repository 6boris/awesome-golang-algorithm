package Solution

import (
	"fmt"
	"testing"
)

func TestIsSameTree(t *testing.T) {
	t.Run("test tree seilize", func(t *testing.T) {
		tree := conFromPreStr("124##5##36##")

		fmt.Println(dumpTreeToString(tree))
	})
}

func TestGenTree(t *testing.T) {
	treedata := []int{3, 9, 20, -1, -1, 15, 7}
	root := GenTree(treedata)
	fmt.Println(root.Val)
	fmt.Println(root.Left.Left)
	fmt.Println(root.Right)
}
