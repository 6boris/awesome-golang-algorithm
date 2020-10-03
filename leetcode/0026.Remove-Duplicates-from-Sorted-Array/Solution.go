package Solution

import "fmt"

func removeDuplicates(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}
	tail := 1
	for i := 1; i < len(nums); i++ {
		if nums[i-1] != nums[i] {
			nums[tail] = nums[i]
			tail++
		}
	}
	fmt.Println(nums)
	return tail
}
