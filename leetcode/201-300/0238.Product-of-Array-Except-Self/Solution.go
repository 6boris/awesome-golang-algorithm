package Solution

func Solution(nums []int) []int {
	length := len(nums)
	r := make([]int, length)
	if length == 0 || length == 1 {
		return r
	}

	r[0] = nums[0]
	for i := 1; i < length-1; i++ {
		r[i] = r[i-1] * nums[i]
	}
	r[length-1] = r[length-2]

	rightProduct := nums[length-1]
	for idx := length - 2; idx > 0; idx-- {
		r[idx] = r[idx-1] * rightProduct
		rightProduct *= nums[idx]
	}
	r[0] = rightProduct

	return r
}
