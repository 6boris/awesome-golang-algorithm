package Solution

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}
type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}
func (this *Codec) preOrder(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	ans := make([]int, 0)
	l := this.preOrder(root.Left)
	r := this.preOrder(root.Right)
	ans = append(ans, root.Val)
	ans = append(ans, l...)
	ans = append(ans, r...)
	return ans
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	if root == nil {
		return ""
	}
	pre := this.preOrder(root)
	dst := make([]int, len(pre))
	copy(dst, pre)
	sort.Ints(dst)
	buf := strings.Builder{}
	for i := 0; i < len(pre); i++ {
		if i != 0 {
			buf.WriteByte(',')
		}
		buf.WriteString(fmt.Sprintf("%d", pre[i]))
	}
	buf.WriteByte('#')
	for i := 0; i < len(dst); i++ {
		if i != 0 {
			buf.WriteByte(',')
		}
		buf.WriteString(fmt.Sprintf("%d", dst[i]))
	}
	return buf.String()
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	if data == "" {
		return nil
	}
	builder := strings.Split(data, "#")
	if len(builder) != 2 {
		return nil
	}
	pre := strings.Split(builder[0], ",")
	in := strings.Split(builder[1], ",")
	return this.buildTree(pre, in)
}

func (this *Codec) buildTree(pre, in []string) *TreeNode {
	if len(pre) == 1 {
		v, _ := strconv.Atoi(pre[0])
		return &TreeNode{Val: v}
	}

	v := pre[0]
	index := 0
	for ; index < len(pre) && in[index] != v; index++ {
	}

	vv, _ := strconv.Atoi(v)
	node := &TreeNode{Val: vv}
	if index > 0 {
		node.Left = this.buildTree(pre[1:index+1], in[:index])
	}
	if index < len(pre)-1 {
		node.Right = this.buildTree(pre[index+1:], in[index+1:])
	}
	return node
}

func Solution(tree *TreeNode) (*TreeNode, string) {
	c := Constructor()
	s := c.serialize(tree)
	t := c.deserialize(s)
	return t, s
}
