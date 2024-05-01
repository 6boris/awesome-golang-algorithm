package Solution

func Solution(nums []int, k int) int {
	l := len(nums)
	sum := 0
	m := 0
	cache := make([][]int, l)
	for i := 0; i < l; i++ {
		sum += nums[i]
		m = max(m, nums[i])
		cache[i] = make([]int, k+1)
		for j := 0; j <= k; j++ {
			cache[i][j] = -1
		}
		cache[i][1] = sum
		if i+1 <= k {
			cache[i][i+1] = m
		}
	}
	var dfs func(int, int) int
	dfs = func(index, kk int) int {
		if index < 0 {
			return -1
		}
		if cache[index][kk] != -1 {
			return cache[index][kk]
		}
		cur := 0
		ans := -1
		for choose := index; choose >= kk-1; choose-- {
			cur += nums[choose]
			if r := dfs(choose-1, kk-1); r != -1 {
				x := max(cur, r)
				if ans == -1 || ans > x {
					ans = x
				}
			}
		}
		cache[index][kk] = ans
		return ans
	}
	return dfs(l-1, k)
}
