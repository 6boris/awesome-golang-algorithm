package Solution

type MultiChildNode struct {
	Val      int
	Children []*MultiChildNode
}

func Solution(root *MultiChildNode) []int {
	ans := make([]int, 0)
	postorder590_(root, &ans)
	return ans
}

func postorder590_(root *MultiChildNode, ans *[]int) {
	if root == nil {
		return
	}
	for _, cNode := range root.Children {
		postorder590_(cNode, ans)
	}
	*ans = append(*ans, root.Val)
}
