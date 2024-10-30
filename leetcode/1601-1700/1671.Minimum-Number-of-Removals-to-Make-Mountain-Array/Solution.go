package Solution

func Solution(nums []int) int {
	l := len(nums)
	if l < 3 {
		return 0
	}
	left, right := make([]int, l), make([]int, l)
	left[0] = 1
	for i := 1; i < l; i++ {
		left[i] = 1
		for pre := i - 1; pre >= 0; pre-- {
			if nums[pre] < nums[i] {
				left[i] = max(left[i], left[pre]+1)
			}
		}
	}
	right[l-1] = 1
	for i := l - 2; i >= 0; i-- {
		right[i] = 1
		for next := i + 1; next < l; next++ {
			if nums[i] > nums[next] {
				right[i] = max(right[i], right[next]+1)
			}
		}
	}
	ans := 0
	for i := 1; i < l-1; i++ {
		if left[i] == 1 || right[i] == 1 {
			continue
		}
		now := left[i] + right[i] - 1
		ans = max(ans, now)
	}
	return l - ans
}
