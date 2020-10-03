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

func moveZeroes2(nums []int) []int {
	n := len(nums)
	if n > 1 {
		for z, nz := 0, 0; nz < n; nz++ {
			if nums[nz] != 0 {
				nums[z], nums[nz] = nums[nz], nums[z]
				z++
			}
		}
	}
	return nums
}
