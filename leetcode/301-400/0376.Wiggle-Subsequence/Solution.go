package Solution

func wiggleMaxLength(nums []int) int {
	//	序列个数小于二直接是摇摆序列
	if len(nums) < 2 {
		return len(nums)
	}
	//	3种状态
	const BEGIN, UP, DOWN = 0, 1, 2
	//	摇摆序列最小长度至少为1
	STATE, maxLength := BEGIN, 1

	for i := 1; i < len(nums); i++ {
		switch STATE {
		case BEGIN:
			if nums[i-1] < nums[i] {
				STATE = UP
				maxLength++
			} else if nums[i-1] > nums[i] {
				STATE = DOWN
				maxLength++
			}
		case UP:
			if nums[i-1] > nums[i] {
				STATE = DOWN
				maxLength++
			}
		case DOWN:
			if nums[i-1] < nums[i] {
				STATE = UP
				maxLength++
			}
		}
	}
	return maxLength
}

func wiggleMaxLength1(nums []int) int {
	length := len(nums)
	if length == 0 {
		return 0
	}

	a, b := 1, 1
	for idx := 1; idx < length; idx++ {
		if nums[idx] > nums[idx-1] {
			a = b + 1
		}
		if nums[idx] < nums[idx-1] {
			b = a + 1
		}

	}
	if a > b {
		return a
	}
	return b
}
