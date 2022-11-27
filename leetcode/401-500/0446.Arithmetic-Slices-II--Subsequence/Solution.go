package Solution

func Solution(nums []int) int {
	l := len(nums)
	if l < 3 {
		return 0
	}
	diffCache := make(map[int]map[int]int)
	diffCache[l-2] = map[int]int{
		nums[l-1] - nums[l-2]: 1,
	}
	ans := 0

	for idx := l - 3; idx >= 0; idx-- {
		diffCache[idx] = make(map[int]int)

		for next := idx + 1; next < l; next++ {
			diff := nums[next] - nums[idx]
			cal := 1

			if cache, ok := diffCache[next]; ok {
				count, ok := cache[diff]
				if ok {
					cal += count
					ans += cal - 1
				}
			}
			diffCache[idx][diff] += cal
		}
	}
	return ans
}
