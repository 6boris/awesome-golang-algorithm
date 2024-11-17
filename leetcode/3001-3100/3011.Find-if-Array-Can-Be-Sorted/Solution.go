package Solution

func countOfOne3011(n int) int {
	c := 0
	for n != 0 {
		c++
		n = n & (n - 1)
	}
	return c
}

func Solution(nums []int) bool {
	ma, mi := nums[0], nums[0]
	cur := countOfOne3011(nums[0])
	list := make([][2]int, 0)
	for i := 1; i < len(nums); i++ {
		now := countOfOne3011(nums[i])
		if now == cur {
			ma = max(ma, nums[i])
			mi = min(mi, nums[i])
			continue
		}
		list = append(list, [2]int{ma, mi})
		cur = now
		ma, mi = nums[i], nums[i]
	}
	list = append(list, [2]int{ma, mi})
	for i := 1; i < len(list); i++ {
		if list[i][1] <= list[i-1][0] {
			return false
		}
	}
	return true
}
