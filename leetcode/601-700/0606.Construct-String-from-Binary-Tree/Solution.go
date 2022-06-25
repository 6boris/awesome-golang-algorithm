package Solution

import (
	"fmt"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func Solution(root *TreeNode) string {
	sb := strings.Builder{}
	buildStringFromTree(root, &sb)
	return sb.String()
}

func buildStringFromTree(root *TreeNode, sb *strings.Builder) {
	if root == nil {
		return
	}
	sb.WriteString(fmt.Sprintf("%d", root.Val))
	if root.Left == nil && root.Right == nil {
		return
	}
	sb.WriteByte('(')
	buildStringFromTree(root.Left, sb)
	sb.WriteByte(')')

	if root.Right != nil {
		sb.WriteByte('(')
		buildStringFromTree(root.Right, sb)
		sb.WriteByte(')')
	}
}
