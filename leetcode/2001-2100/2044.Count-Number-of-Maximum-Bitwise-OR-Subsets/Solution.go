package Solution

func Solution(nums []int) int {
	count := make(map[int]int)
	var dfs func(int, int, int)
	dfs = func(idx, res, l int) {
		if l == 0 {
			count[res]++
			return
		}
		if idx >= len(nums) {
			return
		}

		dfs(idx+1, res|nums[idx], l-1)
		dfs(idx+1, res, l)
	}
	for l := 1; l <= len(nums); l++ {
		dfs(0, 0, l)
	}
	m, ans := -1, 0
	for k, c := range count {
		if m == -1 || k > m {
			m = k
			ans = c
		}
	}
	return ans
}
