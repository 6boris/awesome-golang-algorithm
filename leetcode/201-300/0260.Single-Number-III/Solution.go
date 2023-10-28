package Solution

func Solution(nums []int) []int {
	a := nums[0]
	for i := 1; i < len(nums); i++ {
		a ^= nums[i]
	}
	bitIndex := 0
	for a&1 == 0 {
		bitIndex++
		a >>= 1
	}
	c, d := 0, 0
	for _, n := range nums {
		if (n>>bitIndex)&1 == 1 {
			c ^= n
		} else {
			d ^= n
		}
	}
	return []int{c, d}
}
