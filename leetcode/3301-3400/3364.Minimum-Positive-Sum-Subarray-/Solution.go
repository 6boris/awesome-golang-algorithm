package Solution

import "math"

func Solution(nums []int, l int, r int) int {
	n := len(nums)
	sum := make([]int, n)
	sum[0] = nums[0]
	for i := 1; i < n; i++ {
		sum[i] = sum[i-1] + nums[i]
	}

	ret := math.MaxInt
	for length := l; length <= r; length++ {
		remove := 0
		for e := length - 1; e < n; e++ {
			if s := sum[e] - remove; s > 0 {
				ret = min(ret, s)
			}
			remove = sum[e+1-length]
		}
	}
	if ret == math.MaxInt {
		return -1
	}
	return ret
}
