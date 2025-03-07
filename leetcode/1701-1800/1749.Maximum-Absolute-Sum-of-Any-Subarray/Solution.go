package Solution

func Solution(nums []int) int {
	ans := 0
	_max, _min := 0, 0 //最大的正数，最小的负数
	sum := 0
	for _, n := range nums {
		cur := n
		if cur < 0 {
			cur = -n
		}
		ans = max(ans, cur)
		sum += n
		if sum >= 0 {
			cur = sum - _min
			_max = max(_max, sum)
		} else {
			cur = -(sum - _max)
			_min = min(_min, sum)
		}
		ans = max(ans, cur)
	}
	return ans
}
