package Solution

func Solution(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
    }
	res := make([][]int, 0)
    stack := []*TreeNode{root}
    reverse := false
    for len(stack) > 0 {
        temp := make([]*TreeNode, 0)
        tempVals := make([]int, 0)
        for len(stack) > 0 {
            node := stack[0]
            tempVals = append(tempVals, node.Val)
            if node.Left != nil {
                temp = append(temp, node.Left)
            }
            if node.Right != nil {
                temp = append(temp, node.Right)
            }
            stack = stack[1:]
        }
        if reverse {
            for i, j := 0, len(tempVals) - 1; i < j; i, j = i + 1, j - 1 {
                tempVals[i], tempVals[j] = tempVals[j], tempVals[i]
            }
        }
        reverse = !reverse
        res = append(res, tempVals)
        stack = append(stack, temp...)
    }
    return res
}
