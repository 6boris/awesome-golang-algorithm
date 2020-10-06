package Solution

func twoSum(nums []int, target int) []int {
	left, right := 0, len(nums)-1
	for left < right {
		if nums[left]+nums[right] < target {
			left++
		} else if nums[left]+nums[right] > target {
			right--
		} else {
			return []int{nums[left], nums[right]}
		}
	}
	return nil
}
