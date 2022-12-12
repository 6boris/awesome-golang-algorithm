package Solution

func Solution(nums []int) int {
	max, min := nums[0], nums[0]
	check := make(map[int]int)
	for _, n := range nums {
		if n > max {
			max = n
		}
		if n < min {
			min = n
		}
		check[n]++
	}
	if len(check) > 2 {
		return len(nums) - check[max] - check[min]
	}
	return 0
}
