package Solution

func Solution(nums []int) []int {
	ans := make([]int, len(nums))
	for idx, item := range nums {
		ans[idx] = nums[item]
	}
	return ans
}
