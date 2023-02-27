package Solution

type QuadNode struct {
	Val         bool
	IsLeaf      bool
	TopLeft     *QuadNode
	TopRight    *QuadNode
	BottomLeft  *QuadNode
	BottomRight *QuadNode
}

func Solution(grid [][]int) *QuadNode {
	return construct427(grid, 0, 0, len(grid))
}

func construct427(grid [][]int, topX, topY, length int) *QuadNode {
	node := &QuadNode{}
	if length == 1 {
		node.IsLeaf = true
		node.Val = grid[topX][topY] == 1
		return node
	}

	midLen := length / 2
	tl := construct427(grid, topX, topY, midLen)
	tr := construct427(grid, topX, topY+midLen, midLen)
	bl := construct427(grid, topX+midLen, topY, midLen)
	br := construct427(grid, topX+midLen, topY+midLen, midLen)
	equal := tl.Val == tr.Val && tr.Val == bl.Val && bl.Val == br.Val
	node.Val = true
	if equal {
		node.Val = tl.Val
	}

	node.IsLeaf = tl.IsLeaf && tr.IsLeaf && bl.IsLeaf && br.IsLeaf && equal
	if !node.IsLeaf {
		node.BottomRight = br
		node.BottomLeft = bl
		node.TopLeft = tl
		node.TopRight = tr
	}
	return node
}
