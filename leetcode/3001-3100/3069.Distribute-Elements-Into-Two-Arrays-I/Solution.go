package Solution

func Solution(nums []int) []int {
	size := len(nums)
	ret := make([]int, len(nums))
	ret[0], ret[size-1] = nums[0], nums[1]
	a1Index, a2Index := 0, size-1
	// ret[0], ret[1]
	for i := 2; i < len(nums); i++ {
		if ret[a1Index] > ret[a2Index] {
			a1Index++
			ret[a1Index] = nums[i]
			continue
		}
		a2Index--
		ret[a2Index] = nums[i]
	}
	for s, e := a2Index, size-1; s < e; s, e = s+1, e-1 {
		ret[s], ret[e] = ret[e], ret[s]
	}

	return ret
}
