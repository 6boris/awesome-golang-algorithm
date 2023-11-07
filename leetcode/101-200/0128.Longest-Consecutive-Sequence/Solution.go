package Solution

import "sort"

func longestConsecutive(nums []int) int {
	ans := 0
	m := map[int]int{}
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	for _, v := range nums {
		if _, ok := m[v-1]; !ok {
			m[v] = 1
		} else {
			m[v] = m[v-1] + 1
		}
	}
	for _, v := range m {
		if v > ans {
			ans = v
		}
	}
	return ans
}

func longestConsecutive2(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	sort.Ints(nums)
	ans, tmp := 0, 1
	for idx := 1; idx < len(nums); idx++ {
		r := nums[idx] - nums[idx-1]
		if r == 1 || r == 0 {
			if r == 1 {
				tmp++
			}
			continue
		}
		if tmp > ans {
			ans = tmp
		}
		tmp = 1
	}
	if tmp > ans {
		return tmp
	}
	return ans
}
