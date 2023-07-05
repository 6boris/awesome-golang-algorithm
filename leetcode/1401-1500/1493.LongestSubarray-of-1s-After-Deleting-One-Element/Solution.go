package Solution

func Solution(nums []int) int {
	onlyOneZeroIndices := make([]int, 0)
	l := len(nums)
	zeroCount := 0
	singleMax := 0
	haveZero := false
	for idx := 0; idx < l; idx++ {
		if nums[idx] == 0 {
			haveZero = true
			zeroCount++
			continue
		}
		if zeroCount == 1 {
			onlyOneZeroIndices = append(onlyOneZeroIndices, idx-1)
		}
		zeroCount = 0
		if idx > 0 {
			nums[idx] += nums[idx-1]
		}
		if nums[idx] > singleMax {
			singleMax = nums[idx]
		}
	}
	if len(onlyOneZeroIndices) == 0 {
		if !haveZero {
			singleMax--
		}
		return singleMax
	}

	for _, idx := range onlyOneZeroIndices {
		base := 0
		if idx > 0 {
			base = nums[idx-1]
		}
		next := idx + 1
		for ; next < l && nums[next] > 0; next++ {
		}
		base += nums[next-1]
		if base > singleMax {
			singleMax = base
		}
	}

	return singleMax
}
