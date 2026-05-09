package Solution

func Solution(nums []int) int {
	ret, diff := -1, 0
	size := len(nums)
	prev, cnt := -2, 0
	for i := 1; i < size; i++ {
		diff = nums[i] - nums[i-1]
		if diff == 1 {
			if prev == -2 || prev == -1 {
				cnt++
			} else {
				cnt = 1
			}
			prev = diff
			ret = max(ret, cnt)
			continue
		}
		if diff == -1 {
			if prev == 1 {
				cnt++
				ret = max(ret, cnt)
				prev = diff
				continue
			}
		}
		prev, cnt = -2, 0
	}
	ret = max(ret, cnt)
	if ret == 0 {
		return -1
	}
	return ret + 1
}
