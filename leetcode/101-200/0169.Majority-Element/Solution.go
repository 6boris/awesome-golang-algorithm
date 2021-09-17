package Solution

import "sort"

// Hash
func majorityElement_1(nums []int) int {
	m := make(map[int]int)
	for _, v := range nums {
		m[v]++
		if m[v] > len(nums)/2 {
			return v
		}
	}
	return 0
}

// 排序
func majorityElement_2(nums []int) int {
	sort.Ints(nums)
	return nums[len(nums)/2]
}

//	解法5: 分治算法
func majorityElement_3(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	n := len(nums)

	if n%2 == 1 {
		count := 1

		for i := 0; i < n-1; i++ {
			if nums[n-1] == nums[i] {
				count++
			}
		}
		if count > n/2 {
			return nums[n-1]
		}
	}
	numsTmp := []int{}

	for i := 0; i < n/2; i++ {
		if nums[2*i] == nums[2*i+1] {
			numsTmp = append(numsTmp, nums[2*i])
		}
	}

	return majorityElement_3(numsTmp)
}
