package Solution

func Solution(rolls []int, k int) int {
	// 4, 2, 1, 2, 3, 3, 2, 4, 1
	ans := 1
	cache := make(map[int]struct{})
	for i := 0; i < len(rolls); i++ {
		cache[rolls[i]] = struct{}{}
		if len(cache) == k {
			ans++
			cache = map[int]struct{}{}
		}
	}
	return ans
}
