package Solution

func combinationSum3(k int, n int) (result [][]int) {
	dfs(&result, make([]int, 0), n, 1, k)
	return
}

func dfs(result *[][]int, comb []int, remain int, num int, length int) {
	if remain < 0 || len(comb) > length {
		return
	}
	if remain == 0 && len(comb) == length {
		combCopy := make([]int, len(comb))
		copy(combCopy, comb)
		*result = append(*result, combCopy)
		return
	}

	for i := num; i <= 9; i++ {
		comb = append(comb, i)
		dfs(result, comb, remain-i, i+1, length)
		comb = comb[:len(comb)-1]
	}
}
