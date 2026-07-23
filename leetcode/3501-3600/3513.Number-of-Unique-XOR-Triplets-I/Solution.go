package Solution

func Solution(nums []int) int {
	n := len(nums)
	if n < 3 {
		return n
	}

	mask := 1
	for mask < n {
		mask = (mask << 1) | 1
	}

	return mask + 1
}
