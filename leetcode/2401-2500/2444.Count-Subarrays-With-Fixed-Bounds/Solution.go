package Solution

func Solution(nums []int, minK, maxK int) int64 {
	ans := int64(0)
	preMin, preMax, edge := -1, -1, -1
	for i := 0; i < len(nums); i++ {
		if nums[i] < minK || nums[i] > maxK {
			edge = i
			continue
		}
		if nums[i] == minK {
			preMin = i
		}
		if nums[i] == maxK {
			preMax = i
		}
		which := preMin
		if preMax < which {
			which = preMax
		}
		if diff := which - edge; diff > 0 {
			ans += int64(diff)
		}
	}
	return ans
}
