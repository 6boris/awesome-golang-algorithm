package Solution

func Solution(nums []int) bool {
	l := len(nums)
	incr1, incr2 := 1, l-1
	for ; incr1 < l && nums[incr1] > nums[incr1-1]; incr1++ {
	}
	incr1--
	for ; incr2 > incr1 && nums[incr2-1] < nums[incr2]; incr2-- {

	}
	if incr1 == 0 || incr2 == l-1 || incr1 == incr2 {
		return false
	}
	for ; incr1 < incr2 && nums[incr1] > nums[incr1+1]; incr1++ {

	}
	return incr1 == incr2
}
