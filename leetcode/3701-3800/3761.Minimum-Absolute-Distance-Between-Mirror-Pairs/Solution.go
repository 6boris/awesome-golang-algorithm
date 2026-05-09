package Solution

import "sort"

func reverse3761(n int) int {
	ret := 0
	for n > 0 {
		ret = ret*10 + n%10
		n /= 10
	}
	return ret
}

func Solution(nums []int) int {
	group := make(map[int][]int)
	for i := range nums {
		group[nums[i]] = append(group[nums[i]], i)
	}

	ret := -1
	for i := range nums {
		x := reverse3761(nums[i])

		indies, ok := group[x]
		if !ok {
			continue
		}
		index := sort.Search(len(indies), func(idx int) bool {
			return indies[idx] > i
		})
		if index == len(indies) {
			continue
		}
		maybe := indies[index] - i
		if maybe < 0 {
			maybe = -maybe
		}
		if ret == -1 || ret > maybe {
			ret = maybe
		}
	}

	return ret
}
