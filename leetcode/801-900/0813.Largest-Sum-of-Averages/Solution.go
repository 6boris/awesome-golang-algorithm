package Solution

func Solution(nums []int, k int) float64 {
	l := len(nums)
	cache := make([][]float64, l)
	sum := 0
	mm := 0
	for i := 0; i < l; i++ {
		sum += nums[i]
		mm = max(mm, nums[i])
		cache[i] = make([]float64, k+1)
		for j := 0; j <= k; j++ {
			cache[i][j] = -1.0
		}
		cache[i][1] = float64(sum) / float64(i+1)
		if i+1 <= k {
			cache[i][i+1] = float64(sum)
		}
	}
	var dfs func(int, int) float64
	dfs = func(index, kk int) float64 {
		if index < 0 {
			return 0
		}
		if r := cache[index][kk]; r+1.0 > 0.000001 {
			return r
		}
		ans := float64(0)
		cur := 0
		for c := index; c >= kk-1; c-- {
			cur += nums[c]
			avg := float64(cur) / float64(index-c+1)

			if r := dfs(c-1, kk-1); r+1.0 > 0.000001 {
				ans = max(ans, avg+r)
			}
		}
		cache[index][kk] = ans
		return ans
	}
	return dfs(l-1, k)
}
