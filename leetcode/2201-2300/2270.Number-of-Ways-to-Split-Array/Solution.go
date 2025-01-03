package Solution

func Solution(nums []int) int {
	ans := 0
	l := len(nums)
	for i := 1; i < l; i++ {
		nums[i] += nums[i-1]
	}
	for i := 0; i < l-1; i++ {
		if nums[i] >= nums[l-1]-nums[i] {
			ans++
		}
	}
	return ans
}
