package Solution

func Solution(nums []int) int {
	nega := 0
	for _, n := range nums {
		if n == 0 {
			return 0
		}
		if n < 0 {
			nega++
		}
	}
	if nega&1 == 1 {
		return -1
	}
	return 1
}
