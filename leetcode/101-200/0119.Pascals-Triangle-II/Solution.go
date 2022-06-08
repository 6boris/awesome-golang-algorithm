package Solution

func Solution(rowIndex int) []int {
	if rowIndex == 0 {
		return []int{1}
	}
	pt := make([]int, 0)
	for i := 0; i < rowIndex; i++ {
		tmp := make([]int, 0)
		tmp = append(tmp, 1)
		for j := 1; j < i+1; j++ {
			tmp = append(tmp, pt[j]+pt[j-1])
		}
		tmp = append(tmp, 1)
		pt = tmp
	}
	return pt
}

func Solution2(rowIndex int) []int {
	ans := make([]int, rowIndex+1)
	ans[0] = 1
	for idx := 1; idx <= rowIndex; idx++ {
		for in := idx; in >= 1; in-- {
			ans[in] += ans[in-1]
		}
	}
	return ans
}
