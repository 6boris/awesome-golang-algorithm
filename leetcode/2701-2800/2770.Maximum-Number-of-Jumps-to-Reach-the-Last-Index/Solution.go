package Solution

func Solution(nums []int, target int) int {
	cache := make(map[int]int)
	size := len(nums)
	var dfs func(int) int
	dfs = func(index int) int {
		if index == size-1 {
			return 0
		}
		if v, ok := cache[index]; ok {
			return v
		}
		ret := -1
		for next := index + 1; next < size; next++ {
			cmp := nums[next] - nums[index]
			if cmp >= -target && cmp <= target {
				if test := dfs(next); test != -1 {
					ret = max(ret, test+1)
				}
			}
		}
		cache[index] = ret
		return ret
	}
	return dfs(0)
}
