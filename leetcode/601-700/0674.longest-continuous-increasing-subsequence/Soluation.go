package Soluation

func findLengthOfLCIS(nums []int) int {
	if 0 == len(nums) {
		return 0
	}
	ans, sum := 1, 1
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			sum++
		} else {
			sum = 1
		}
		if sum > ans {
			ans = sum
		}
	}
	return ans
}
