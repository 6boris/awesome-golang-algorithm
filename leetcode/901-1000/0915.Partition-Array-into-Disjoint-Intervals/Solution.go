package Solution

func Solution(nums []int) int {
	length := len(nums)
	max, min := make([]int, length), make([]int, length)
	max[0], min[length-1] = nums[0], nums[length-1]
	for idx := 1; idx < length; idx++ {
		max[idx] = max[idx-1]
		if nums[idx] > max[idx] {
			max[idx] = nums[idx]
		}

		i := length - 1 - idx
		min[i] = min[i+1]
		if nums[i] < min[i] {
			min[i] = nums[i]
		}
	}

	idx := 1
	for ; idx < length; idx++ {
		if min[idx] >= max[idx-1] {
			break
		}
	}

	return idx
}
