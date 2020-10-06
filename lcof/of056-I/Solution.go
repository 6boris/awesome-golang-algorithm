package Solution

func singleNumbers(nums []int) []int {
	tmp := 0
	for _, num := range nums {
		tmp ^= num
	}
	div := 1

	for div&tmp == 0 {
		div <<= 1
	}
	a, b := 0, 0

	for _, num := range nums {
		if div&num == 0 {
			a ^= num
		} else {
			b ^= num
		}
	}
	return []int{a, b}
}
