package Solution

import "sort"

func lengthOfLIS(nums []int) int {
	tails := make([]int, len(nums)) // 长度为i的子数组最小的末尾数组
	var result int
	for x := 0; x < len(nums); x++ {
		var i, j int = 0, result
		for i != j {
			m := (i + j) / 2
			if tails[m] < nums[x] {
				i = m + 1
			} else {
				j = m
			}
		}
		tails[i] = nums[x]
		if i == result {
			result++
		}
	}
	return result
}

func lengthOfLIS2(nums []int) int {
	t := []int{}
	for _, n := range nums {
		j := sort.Search(len(t), func(i int) bool {
			return t[i] >= n
		})
		if j == len(t) {
			t = append(t, n)
		} else {
			t[j] = n
		}
	}
	return len(t)
}
