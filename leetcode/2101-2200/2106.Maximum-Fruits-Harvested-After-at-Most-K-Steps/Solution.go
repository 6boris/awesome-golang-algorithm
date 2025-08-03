package Solution

import "sort"

func Solution(fruits [][]int, startPos int, k int) int {
	l := len(fruits)
	index := sort.Search(l, func(i int) bool {
		return fruits[i][0] >= startPos
	})
	ans, dis := 0, 0
	sum := make([]int, l)

	a := 0
	for i := index; i < l; i++ {
		a += fruits[i][1]
		sum[i] = a
	}
	a = 0
	for i := index - 1; i >= 0; i-- {
		a += fruits[i][1]
		sum[i] = a
	}
	right := 0
	for i := index; i < l; i++ {
		dis = fruits[i][0] - startPos
		if dis > k {
			break
		}
		right += fruits[i][1]
		ans = max(ans, right)

		if x := sort.Search(index, func(j int) bool {
			return fruits[i][0]-fruits[j][0] <= k-dis
		}); x != index {
			ans = max(ans, right+sum[x])
		}
	}
	left := 0
	ll := l - index

	for i := index - 1; i >= 0; i-- {
		dis = startPos - fruits[i][0]
		if dis > k {
			break
		}
		left += fruits[i][1]

		ans = max(ans, left)
		if x := sort.Search(ll, func(j int) bool {
			return fruits[index+j][0]-fruits[i][0] > k-dis
		}); x != 0 {
			ans = max(ans, left+sum[index+x-1])

		}
	}
	return ans
}
