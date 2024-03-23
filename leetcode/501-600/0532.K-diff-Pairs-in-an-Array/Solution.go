package Solution

import "sort"

func Solution(nums []int, k int) int {
	n := len(nums)
	if n < 2 || k < 0 {
		return 0
	}
	hash, c := make(map[int]int), 0
	for _, num := range nums {
		hash[num]++
	}
	if k == 0 {
		for _, v := range hash {
			if v >= 2 {
				c++
			}
		}
	} else {
		for key := range hash {
			if hash[key+k] > 0 {
				c++
			}
		}
	}
	return c
}

func Solution1(nums []int, k int) int {
	sort.Ints(nums)
	ans := 0
	for i := 0; i < len(nums)-1; i++ {
		if i != 0 && nums[i] == nums[i-1] {
			continue
		}
		idx := sort.Search(len(nums), func(ii int) bool {
			return ii > i && nums[ii] >= nums[i]+k
		})
		if idx != len(nums) && nums[idx] == nums[i]+k {
			ans++
		}
	}
	return ans
}
