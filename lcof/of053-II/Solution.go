package Solution

import (
	"sort"
)

func missingNumber_1(nums []int) int {
	sort.Ints(nums)
	if nums[0] != 0 {
		return 0
	} else if nums[len(nums)-1] != len(nums) {
		return len(nums)
	}
	for i := 1; i < len(nums); i++ {
		if i != nums[i] {
			return i
		}
	}
	return 0
}

func missingNumber_2(nums []int) int {
	m := make(map[int]int)
	for _, v := range nums {
		m[v]++
	}
	for i := range nums {
		if _, ok := m[i+1]; !ok {
			return i + 1
		}
	}
	return 0
}

func missingNumber_3(nums []int) int {
	ans := len(nums)
	for i, v := range nums {
		ans ^= i ^ v
	}
	return ans
}

func missingNumber_4(nums []int) int {
	sum := len(nums) * (len(nums) + 1) / 2
	for _, v := range nums {
		sum -= v
	}
	return sum
}
