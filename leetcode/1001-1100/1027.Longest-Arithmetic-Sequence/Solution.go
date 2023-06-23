package Solution

func Solution(nums []int) int {
	ans := 0
	cache := make(map[int]map[int]int)
	for i := len(nums) - 2; i >= 0; i-- {
		cache[i] = make(map[int]int)
		for next := i + 1; next < len(nums); next++ {
			diff := nums[next] - nums[i]
			add := 2
			v, ok := cache[next][diff]
			if ok {
				add = 1
			}
			add += v
			if v, ok := cache[i][diff]; !ok || v < add {
				cache[i][diff] = add
			}
		}
		for _, v := range cache[i] {
			if v > ans {
				ans = v
			}
		}
	}
	return ans
}
