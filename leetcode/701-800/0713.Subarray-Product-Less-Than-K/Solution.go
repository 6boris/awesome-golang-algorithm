package Solution

func Solution(nums []int, k int) int {
	cache := make(map[int]int)
	ans := 0
	if nums[0] < k {
		cache[nums[0]] = 1
		ans++
	}
	for i := 1; i < len(nums); i++ {
		if nums[i] >= k {
			cache = map[int]int{}
			continue
		}
		next := map[int]int{nums[i]: 1}
		for kk, v := range cache {
			if r := nums[i] * kk; r < k {
				next[r] += v
			}
		}
		cache = next
		for _, v := range cache {
			ans += v
		}
	}

	return ans
}
