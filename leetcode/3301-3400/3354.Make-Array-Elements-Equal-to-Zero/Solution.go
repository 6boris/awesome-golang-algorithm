package Solution

func Solution(nums []int) int {
	l := len(nums)
	left, right := make([]int, l+2), make([]int, l+2)
	for i := 1; i <= l; i++ {
		left[i] = left[i-1] + nums[i-1]
	}
	for i := l; i >= 1; i-- {
		right[i] = right[i+1] + nums[i-1]
	}
	var ret int
	for index := 1; index <= l; index++ {
		if nums[index-1] != 0 {
			continue
		}
		diff := left[index-1] - right[index+1]
		if diff == 0 {
			ret += 2
		}
		if diff == 1 || diff == -1 {
			ret++
		}
	}
	return ret
}
