package Solution

type MultiChildNode struct {
	Val      int
	Children []*MultiChildNode
}

func Solution(root *MultiChildNode) []int {
	ans := make([]int, 0)
	preorder589_(root, &ans)
	return ans
}

func preorder589_(root *MultiChildNode, ans *[]int) {
	if root == nil {
		return
	}

	*ans = append(*ans, root.Val)
	for _, cNode := range root.Children {
		preorder589_(cNode, ans)
	}
}
