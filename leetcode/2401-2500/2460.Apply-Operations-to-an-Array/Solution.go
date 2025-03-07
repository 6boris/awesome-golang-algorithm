package Solution

func Solution(nums []int) []int {
	ans := make([]int, len(nums))
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == nums[i+1] {
			nums[i], nums[i+1] = nums[i]*2, 0
			continue
		}
	}
	index := 0
	for i := range nums {
		if nums[i] != 0 {
			ans[index] = nums[i]
			index++
		}
	}
	return ans
}
