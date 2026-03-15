package Solution

import "sort"

func Solution(nums []int, k int) bool {
	sum := 0
	testMaxNum := 0
	for i := range nums {
		sum += nums[i]
		testMaxNum = max(testMaxNum, nums[i])
	}
	mod := sum % k
	if mod != 0 || testMaxNum > sum/k {
		return false
	}

	sort.Slice(nums, func(i, j int) bool {
		return nums[i] > nums[j]
	})
	target := sum / k
	buckets := make([]int, k)
	var dfs func(int) bool
	dfs = func(index int) bool {
		if index == len(nums) {
			return true
		}
		for i := range k {
			if i > 0 && buckets[i] == buckets[i-1] {
				// 如果这俩桶一样，前一个放不进去，这个也进不去:wq
				continue
			}
			// 尝试添加到桶里。
			maybe := buckets[i] + nums[index]
			if maybe <= target {
				buckets[i] = maybe
				if dfs(index + 1) {
					return true
				}
				buckets[i] -= nums[index]
			}
			// 空的都放不进去，就不用了
			if buckets[i] == 0 {
				break
			}
		}
		return false
	}
	return dfs(0)
}
