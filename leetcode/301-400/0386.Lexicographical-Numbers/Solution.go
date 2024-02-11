package Solution

func Solution(n int) []int {
	ans := make([]int, 0)
	var dfs func(int, int)
	dfs = func(i int, j int) {
		cur := j*10 + i
		if cur > n {
			return
		}
		ans = append(ans, cur)
		for k := 0; k < 10; k++ {
			dfs(k, cur)
		}
	}
	for i := 1; i < 10; i++ {
		dfs(i, 0)
	}
	return ans
}
