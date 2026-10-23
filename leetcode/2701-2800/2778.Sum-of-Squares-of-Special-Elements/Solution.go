package Solution

func Solution(nums []int) int {
	var ret int
	n := len(nums)
	for i := 1; i <= n; i++ {
		if n%i == 0 {
			ret += nums[i-1] * nums[i-1]
		}
	}
	return ret
}
