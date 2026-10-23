package Solution

func Solution(nums []int, k int) []int64 {
	ret := make([]int64, k)
	if k <= 0 {
		return ret
	}

	if k == 1 {
		n := int64(len(nums))
		total := n * (n + 1) / 2
		ret[0] = total
		return ret
	}

	// curr[rem] 记录以当前元素结尾的连续子数组中，乘积 % k == rem 的子数组数量
	curr := make([]int64, k)

	for _, val := range nums {
		next := make([]int64, k)
		v := val % k
		if v < 0 {
			v += k
		}

		next[v]++

		for rem := 0; rem < k; rem++ {
			if curr[rem] > 0 {
				newRem := (rem * v) % k
				next[newRem] += curr[rem]
			}
		}

		for x := 0; x < k; x++ {
			ret[x] += next[x]
		}

		// 滚动更新
		curr = next
	}

	return ret

}
