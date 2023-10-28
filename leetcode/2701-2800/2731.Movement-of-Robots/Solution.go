package Solution

import (
	"sort"
)

const mod2731 = 1000000007

func Solution(nums []int, s string, d int) int {

	length := len(nums)
	for i := range nums {
		if s[i] == 'R' {
			nums[i] += d
			continue
		}
		nums[i] -= d
	}
	sort.Ints(nums)

	sum := make([][]int, length)
	for i := len(nums) - 1; i >= 0; i-- {
		sum[i] = make([]int, 4)
		if i == len(nums)-1 {
			if nums[i] < 0 {
				sum[i][0] = nums[i]
				sum[i][1] = 1
			} else {
				sum[i][2] = nums[i]
				sum[i][3] = 1
			}
			continue
		}
		if nums[i] < 0 {
			sum[i][0] = sum[i+1][0] + nums[i]
			sum[i][1] = sum[i+1][1] + 1
			sum[i][2] = sum[i+1][2]
			sum[i][3] = sum[i+1][3]
		} else {
			sum[i][2] = sum[i+1][2] + nums[i]
			sum[i][3] = sum[i+1][3] + 1
			sum[i][0] = sum[i+1][0]
			sum[i][1] = sum[i+1][1]
		}
	}
	ans := 0
	var a, b int
	for i := 0; i < length-1; i++ {
		a = nums[i]*sum[i+1][3] - sum[i+1][2]
		b = nums[i]*sum[i+1][1] - sum[i+1][0]
		if a < 0 {
			a = -a
		}
		if b < 0 {
			b = -b
		}
		a += b

		ans = (ans + a) % mod2731
	}
	return ans
}
