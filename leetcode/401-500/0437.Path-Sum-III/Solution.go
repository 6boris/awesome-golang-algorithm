package Solution

func Solution(root *TreeNode, target int) int {
	subSum := make(map[int]int)
	ans := 0
	down2top(root, target, &ans, subSum)
	return ans
}

func Solution2(root *TreeNode, target int) int {
	subSum := make(map[int]int)
	ans := 0
	subSum[0] = 1 // target = target
	top2down(root, target, 0, &ans, subSum)
	return ans
}

func down2top(root *TreeNode, target int, ans *int, subTreeSum map[int]int) {
	if root == nil {
		return
	}

	if root.Left == nil && root.Right == nil {
		subTreeSum[root.Val]++
		if root.Val == target {
			*ans++
		}
		return
	}

	expectSum := make(map[int]int)
	down2top(root.Left, target, ans, expectSum)
	down2top(root.Right, target, ans, expectSum)
	subTreeSum[root.Val]++
	if root.Val == target {
		*ans++
	}
	for k, v := range expectSum {
		subTreeSum[k+root.Val] += v
		if k+root.Val == target {
			*ans += v
		}
	}
}

func top2down(root *TreeNode, target, sum int, ans *int, subTreeSum map[int]int) {
	if root == nil {
		return
	}

	sum += root.Val
	if count, ok := subTreeSum[sum-target]; ok {
		*ans += count
	}
	subTreeSum[sum]++

	top2down(root.Left, target, sum, ans, subTreeSum)
	top2down(root.Right, target, sum, ans, subTreeSum)
	subTreeSum[sum]--
}
