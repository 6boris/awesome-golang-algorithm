package Solution

func Solution(nums []int) []int64 {
	size := len(nums)
	ret := make([]int64, size)

	group := make(map[int][]int64)
	for i, v := range nums {
		if _, ok := group[v]; !ok {
			group[v] = []int64{}
		}
		pick := int64(0)
		n := len(group[v])
		if n > 0 {
			pick = group[v][n-1]
		}
		pick += int64(i)
		group[v] = append(group[v], pick)
	}

	curr := make(map[int]int)

	for i, v := range nums {
		item := group[v]
		n := len(item)
		if n == 1 {
			continue
		}

		idx := curr[v]
		curr[v]++

		x := int64(i)

		var smaller int64
		if idx > 0 {
			smaller = int64(idx)*x - item[idx-1]
		}
		bigger := (item[n-1] - item[idx]) - int64(n-1-idx)*x
		ret[i] = smaller + bigger
	}

	return ret
}
