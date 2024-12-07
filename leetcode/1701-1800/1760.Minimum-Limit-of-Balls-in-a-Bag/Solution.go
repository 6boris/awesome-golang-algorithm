package Solution

func Solution(nums []int, maxOperations int) int {
	var can func(int) bool
	can = func(limit int) bool {
		op := maxOperations
		for _, n := range nums {
			if n <= limit {
				continue
			}
			// 9 /3
			times := n / limit
			if n%limit == 0 {
				times--
			}
			if times > op {
				return false
			}
			op -= times
		}
		return true
	}
	left, right := 1, 1000000001
	ans := -1
	for left < right {
		mid := (left + right) / 2
		if can(mid) {
			right = mid
			ans = mid
			continue
		}
		left = mid + 1
	}
	return ans
}
