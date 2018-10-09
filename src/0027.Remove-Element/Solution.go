package Solution

func removeElement(nums []int, val int) int {
	tail := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[tail] = nums[i]
			tail++
		}
	}
	return tail
}
