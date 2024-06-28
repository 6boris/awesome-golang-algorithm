package Solution

import "sort"

func Solution(n int, roads [][]int) int64 {
	ans := int64(0)

	// 检查入读和出度
	in := make([]int, n)
	for _, r := range roads {
		in[r[0]]++
		in[r[1]]++
	}
	sort.Ints(in)
	v := int64(1)
	for _, i := range in {
		ans += v * int64(i)
		v++
	}
	return ans
}
