package Solution

func gcd1979(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd1979(b, a%b)
}

func Solution(nums []int) int {
	max, min := nums[0], nums[0]
	for idx := 1; idx < len(nums); idx++ {
		if nums[idx] > max {
			max = nums[idx]
		}
		if nums[idx] < min {
			min = nums[idx]
		}
	}
	return gcd1979(max, min)
}
