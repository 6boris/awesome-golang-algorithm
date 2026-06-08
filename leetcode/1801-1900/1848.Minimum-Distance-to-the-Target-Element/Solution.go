package Solution

func Solution(nums []int, target, start int) int {
	ret := -1
	diff := 0
	for i := range nums {
		if nums[i] == target {
			diff = i - start
			if diff < .0 {
				diff = -diff
			}
			if ret == -1 || ret > diff {
				ret = diff
			}
		}
	}
	return ret
}
