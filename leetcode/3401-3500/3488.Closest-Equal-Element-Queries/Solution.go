package Solution

import "sort"

func Solution(nums, queries []int) []int {
	ret := make([]int, len(queries))
	group := make(map[int][]int)
	// 最后一个元素可以
	size := len(nums)
	for i := range nums {
		group[nums[i]] = append(group[nums[i]], i)
	}
	cache := make(map[int]int)
	for q := range queries {
		if v, ok := cache[queries[q]]; ok {
			ret[q] = v
			continue
		}

		test := group[nums[queries[q]]]
		if len(test) <= 1 {
			ret[q] = -1
			cache[queries[q]] = ret[q]
			continue
		}
		index := sort.Search(len(test), func(i int) bool {
			return test[i] >= queries[q]
		})
		if index == 0 {
			ret[q] = min(test[index+1]-test[index], size-test[len(test)-1]+test[0])
		} else if index == len(test)-1 {
			ret[q] = min(test[index]-test[index-1], size-test[len(test)-1]+test[0])
		} else {
			ret[q] = min(test[index]-test[index-1], test[index+1]-test[index])
		}
		cache[queries[q]] = ret[q]
	}

	return ret
}
