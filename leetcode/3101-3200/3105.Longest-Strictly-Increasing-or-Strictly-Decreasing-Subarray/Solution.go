package Solution

func helper(nums []int, less func(i, j int) bool) int {
	l := 1
	res := 1
	for i := 1; i < len(nums); i++ {
		if less(nums[i-1], nums[i]) {
			l++
			res = max(res, l)
			continue
		}
		l = 1
	}
	return max(res, l)
}

func Solution(nums []int) int {
	a := helper(nums, func(i, j int) bool {
		return i < j
	})
	b := helper(nums, func(i, j int) bool {
		return i > j
	})
	return max(a, b)
}
