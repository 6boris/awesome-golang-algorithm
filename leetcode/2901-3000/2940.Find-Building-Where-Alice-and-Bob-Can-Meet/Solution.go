package Solution

import "sort"

func Solution(heights []int, queries [][]int) []int {
	// monotonic stack + sort + binarysearch
	ll := len(queries)
	indies := make([]int, ll)
	for i := range ll {
		if queries[i][0] > queries[i][1] {
			queries[i][0], queries[i][1] = queries[i][1], queries[i][0]
		}
		indies[i] = i
	}
	sort.Slice(indies, func(i, j int) bool {
		return queries[indies[i]][1] > queries[indies[j]][1]
	})

	ans := make([]int, ll)
	for i := range ll {
		ans[i] = -1
	}

	l := len(heights)
	stack := make([]int, l)
	stackIndex := l - 1
	stack[stackIndex] = l - 1
	cur := l - 2

	for _, i := range indies {
		a, b := queries[i][0], queries[i][1]
		if a == b || heights[a] < heights[b] {
			ans[i] = b
			continue
		}
		if b == l-1 {
			continue
		}
		target := max(heights[a], heights[b])
		for ; cur > b; cur-- {
			for ; stackIndex != l && heights[cur] >= heights[stack[stackIndex]]; stackIndex++ {
			}
			stackIndex--
			stack[stackIndex] = cur
		}
		if idx := sort.Search(l-stackIndex, func(i int) bool {
			return heights[stack[i+stackIndex]] > target
		}); idx != l-stackIndex {
			ans[i] = stack[idx+stackIndex]
		}
	}

	return ans
}
