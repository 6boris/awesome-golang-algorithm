package Solution

import (
	"sort"
)

type tmp2191 struct {
	v, s, i int
}

func Solution(mapping []int, nums []int) []int {
	convert := func(n int) int {
		if n == 0 {
			return mapping[0]
		}
		base := 1
		v := 0
		for n > 0 {
			mod := n % 10
			n /= 10
			v = v + base*mapping[mod]
			base *= 10
		}
		return v
	}

	list := make([]tmp2191, len(nums))
	for i, n := range nums {
		list[i] = tmp2191{
			s: n, i: i, v: convert(n),
		}
	}
	sort.Slice(list, func(i, j int) bool {
		if list[i].v == list[j].v {
			return list[i].i < list[j].i
		}
		return list[i].v < list[j].v
	})

	ans := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		ans[i] = list[i].s
	}
	return ans
}
