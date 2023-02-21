package Solution

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func Solution(root *TreeNode) []int {
	sum := make(map[int]int)
	t(root, sum)
	aggregate := make(map[int]int)
	mostFreq, ans := -1, []int{}
	for v, count := range sum {
		aggregate[count]++
		if count > mostFreq {
			ans = []int{v}
			mostFreq = count
		} else if count == mostFreq {
			ans = append(ans, v)
		}
	}
	return ans
}

func t(root *TreeNode, sum map[int]int) int {
	if root == nil {
		return 0
	}
	x := t(root.Left, sum) + t(root.Right, sum) + root.Val
	sum[x]++
	return x
}
