package Solution

func Solution(nums []int) bool {
	c := 0
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] > nums[i+1] {
			if i > 0 {
				if nums[i-1] <= nums[i+1] {
					nums[i] = nums[i-1]
				} else {
					nums[i+1] = nums[i]
				}
			}
			c++
		}
	}
	return c <= 1
}
