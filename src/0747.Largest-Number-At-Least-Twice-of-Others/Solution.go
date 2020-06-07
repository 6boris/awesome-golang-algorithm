package Solution

func Solution(nums []int) int {
	n := len(nums)
	if n == 1 {
		return 0
	}
	max, maxIndex := -1, -1
	for i, num := range nums {
		if num > max {
			max, maxIndex = num, i
		}
	}
	for _, num := range nums {
		if num != max && 2*num > max {
			return -1
		}
	}
	return maxIndex
}
