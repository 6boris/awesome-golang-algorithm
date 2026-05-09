package Solution

import "sort"

const mod1712 = 1000000007

func Solution(nums []int) int {
	n := len(nums)
	sum := make([]int, n)
	sum[0] = nums[0]
	for i := 1; i < n; i++ {
		sum[i] = sum[i-1] + nums[i]
	}

	ret := 0
	for i := 0; i < n-2; i++ {
		leftSum := sum[i]

		if leftSum*3 > sum[n-1] {
			break
		}
		// sum[j] - left >= left
		L := sort.Search(n-1-(i+1), func(k int) bool {
			j := i + 1 + k
			return sum[j] >= 2*leftSum
		}) + i + 1

		// !(sum[j] - left <= sum[n-1] - sum[j])
		R := sort.Search(n-1-(i+1), func(k int) bool {
			j := i + 1 + k
			return 2*sum[j] > sum[n-1]+leftSum
		}) + i + 1

		if R > L {
			ret = (ret + (R - L)) % mod1712
		}
	}

	return ret
}
