package Solution

func Solution(nums []int) []int {
	inQuickSort(nums, 0, len(nums)-1, less)
	return nums
}

func inQuickSort(nums []int, start, end int, less func(i, j int) bool) {
	if r := end - start + 1; r == 0 || r == 1 {
		return
	}
	cmpObj := nums[start]
	lessIdx := start
	for idx := start + 1; idx <= end; idx++ {
		if less(nums[idx], cmpObj) {
			lessIdx++
			nums[lessIdx], nums[idx] = nums[idx], nums[lessIdx]
		}
	}
	nums[start], nums[lessIdx] = nums[lessIdx], nums[start]
	inQuickSort(nums, start, lessIdx-1, less)
	inQuickSort(nums, lessIdx+1, end, less)
}

func less(i, j int) bool {
	iEven := i&1 == 0
	jEven := j&1 == 0
	if iEven && jEven || !iEven && !jEven {
		return i < j
	}
	if !jEven {
		return true
	}
	return false
}

func Solution1(nums []int) []int {
	s, e := 0, len(nums)-1
	for s < e {
		for ; s < e && nums[s]&1 == 0; s++ {
		}
		for ; s < e && nums[e]&1 == 1; e-- {
		}
		if s < e {
			nums[s], nums[e] = nums[e], nums[s]
		}
	}
	return nums
}
