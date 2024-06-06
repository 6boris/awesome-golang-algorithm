package Solution

import "sort"

func Solution(envelopes [][]int) int {
	// 先根据宽度排序，然后根据高度排序
	sort.Slice(envelopes, func(i, j int) bool {
		if envelopes[i][0] == envelopes[j][0] {
			//return envelopes[i][1] < envelopes[j][1]
			// 2, 4
			// 1  1
			// 这情况，导致成两个，实际宽度相同，只能保留一个
			return envelopes[i][1] > envelopes[j][1]

		}
		return envelopes[i][0] < envelopes[j][0]
	})

	// 2, 4, 1, 3, 5
	// 1  1, 2, 2, 3
	// 1, 4
	// 这个问题就是，最长公共子序列问题了
	lis := make([]int, 0)
	for _, h := range envelopes {
		index := sort.Search(len(lis), func(i int) bool {
			return lis[i] >= h[1]
		})
		if index == len(lis) {
			//找不到
			lis = append(lis, h[1])
		} else {
			lis[index] = h[1]
		}
	}

	return len(lis)
}
