package Solution

func Solution(nums []int) int {
	psum, min := 0, 1
	for _, val := range nums {
		psum += val
		if psum < min {
			min = psum
		}
	}
	if min > 0 {
		return min
	}
	return -min + 1
}
