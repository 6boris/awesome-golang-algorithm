package Solution

import "math"

func Solution(nums []int) int64 {
	n := len(nums)
	var p, q, j int
	var max_sum, sum, res int64
	ans := int64(math.MinInt64)
	for i := 0; i < n; i++ {
		j = i + 1
		res = 0
		// first segment
		for ; j < n && nums[j-1] < nums[j]; j++ {
		}
		p = j - 1
		if p == i {
			continue
		}
		// second segment
		res += int64(nums[p] + nums[p-1])
		for ; j < n && nums[j-1] > nums[j]; j++ {
			res += int64(nums[j])
		}
		q = j - 1
		if q == p || q == n-1 || (j < n && nums[j] <= nums[q]) {
			i = q
			continue
		}
		// third segment
		res += int64(nums[q+1])
		// find the maximum sum of the third segment
		max_sum = 0
		sum = 0
		for k := q + 2; k < n && nums[k] > nums[k-1]; k++ {
			sum += int64(nums[k])
			if sum > max_sum {
				max_sum = sum
			}
		}
		res += max_sum
		// find the maximum sum of the first segment
		max_sum = 0
		sum = 0
		for k := p - 2; k >= i; k-- {
			sum += int64(nums[k])
			if sum > max_sum {
				max_sum = sum
			}
		}
		res += max_sum
		// update answer
		if res > ans {
			ans = res
		}
		i = q - 1
	}
	return ans
}
