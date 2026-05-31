package Solution

func Solution(nums []int) int {
	digitSum := func(n int) int {
		sum := 0
		for n > 0 {
			sum += n % 10
			n /= 10
		}
		return sum
	}
	ret := digitSum(nums[0])
	for i := 1; i < len(nums); i++ {
		ret = min(ret, digitSum(nums[i]))
	}
	return ret
}
