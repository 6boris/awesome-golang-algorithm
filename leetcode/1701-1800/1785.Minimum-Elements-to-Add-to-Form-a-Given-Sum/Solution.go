package Solution

func Solution(nums []int, limit int, goal int) int {
	diff := goal
	for _, n := range nums {
		diff -= n
	}
	if diff == 0 {
		return 0
	}
	if diff < 0 {
		diff = -diff
	}
	a := diff / limit
	if diff%limit != 0 {
		a++
	}
	return a
}
