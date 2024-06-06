package Solution

type Node struct {
	Val      int
	Children []*Node
}

func Solution(root *Node) int {
	if root == nil {
		return 0
	}
	dep := 0
	for _, c := range root.Children {
		dep = max(dep, Solution(c))
	}
	return dep + 1
}
