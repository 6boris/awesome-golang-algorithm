package Solution

func Solution(nums []int) []int {
	l := len(nums)
	if l == 1 {
		return []int{-1}
	}

	ans := make([]int, l)

	preStack := make([][]int, l-1)
	preStack[0] = []int{nums[0]}
	for i := 1; i < l-1; i++ {
		if nums[i] > nums[i-1] {
			preStack[i] = append(preStack[i-1], nums[i])
		} else {
			preStack[i] = preStack[i-1]
		}
	}
	stackIndex := l - 1

	stack := make([]int, len(nums))
	copy(stack, nums)
	for idx := l - 1; idx >= 0; idx-- {
		for stackIndex < l && stack[stackIndex] <= nums[idx] {
			stackIndex++
		}
		if stackIndex == l {
			// empty
			set := -1
			if idx != 0 {
				for _, n := range preStack[idx-1] {
					if n > nums[idx] {
						set = n
						break
					}
				}
			}
			ans[idx] = set
		} else {
			ans[idx] = stack[stackIndex]
		}
		stackIndex--
		stack[stackIndex] = nums[idx]
	}
	return ans
}
