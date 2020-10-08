package Solution

func Solution(nums []int) int {
	n := len(nums)
	if n < 1 {
		return -1
	}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	tsum := 0
	for i, num := range nums {
		sum -= num
		if tsum == sum {
			return i
		}
		tsum += num
	}
	return -1
}
