package Solution

func moveZeroes(nums []int) []int {
	count := 0
	for _, i := range nums {
		if i != 0 {
			nums[count] = i
			count++
		}
	}
	for i := count; i < len(nums); i++ {
		temp := i
		nums[temp] = 0
	}
	return nums
}
