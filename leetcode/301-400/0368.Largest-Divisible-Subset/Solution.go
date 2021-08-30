package Solution

import "sort"

func Solution(nums []int) []int {
	length := len(nums)
	if length == 0 {
		return []int{}
	}

	sort.Ints(nums)
	r := make([]int, length)
	// 找回溯数组
	maxidx := 0
	preIdx := make([]int, length)
	for i := 0; i < length; i++ {
		r[i] = 1
		preIdx[i] = i
	}
	for i := 1; i < length; i++ {
		for j := 0; j < i; j++ {
			if nums[i]%nums[j] == 0 && r[j]+1 > r[i] {
				r[i] = r[j] + 1
				preIdx[i] = j
			}
		}
		if r[i] > r[maxidx] {
			maxidx = i
		}
	}
	result := make([]int, 0)
	for preIdx[maxidx] != maxidx {
		result = append(result, nums[maxidx])
		maxidx = preIdx[maxidx]
	}
	result = append(result, nums[maxidx])
	for s, e := 0, len(result)-1; s < e; s, e = s+1, e-1 {
		result[s], result[e] = result[e], result[s]
	}
	return result
}
