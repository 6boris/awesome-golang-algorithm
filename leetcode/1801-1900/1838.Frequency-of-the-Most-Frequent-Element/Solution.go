package Solution

import "sort"

func Solution(nums []int, k int) int {
	sort.Ints(nums)
	prefixSum := make([]int64, len(nums))
	prefixSum[0] = int64(nums[0])
	count := make(map[int]int)
	count[nums[0]]++
	for i := 1; i < len(nums); i++ {
		prefixSum[i] = int64(nums[i]) + prefixSum[i-1]
		count[nums[i]]++
	}
	ans := count[nums[0]]
	v := make(map[int]struct{})
	v[nums[0]] = struct{}{}

	for idx := 1; idx < len(nums); idx++ {
		if _, ok := v[nums[idx]]; ok {
			continue
		}
		v[nums[idx]] = struct{}{}
		i := sort.Search(idx, func(i int) bool {
			sum := int64(nums[idx]) * int64(idx-i)
			psum := int64(0)
			if i == 0 {
				psum = prefixSum[idx-1]
			} else {
				psum = prefixSum[idx-1] - prefixSum[i-1]
			}
			diff := sum - psum
			return diff <= int64(k)
		})
		if r := idx - i + count[nums[idx]]; r > ans {
			ans = r
		}
	}
	return ans
}
