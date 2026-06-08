package Solution

import (
	"strconv"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Codec struct{}

func Constructor() Codec {
	return Codec{}
}

func dfs(valsPtr *[]string) *TreeNode {
	val := (*valsPtr)[0]
	*valsPtr = (*valsPtr)[1:]
	if val == "nil" {
		return nil
	}
	valInt, _ := strconv.Atoi(val)
	node := &TreeNode{Val: valInt}
	node.Left = dfs(valsPtr)
	node.Right = dfs(valsPtr)
	return node
}

func (c *Codec) deserialize(data string) *TreeNode {
	vals := strings.Split(data, ",")
	return dfs(&vals)
}
