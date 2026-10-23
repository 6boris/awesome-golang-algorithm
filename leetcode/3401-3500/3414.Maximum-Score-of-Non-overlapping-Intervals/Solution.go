package Solution

import "sort"

type Interval struct {
	id     int
	l      int
	r      int
	weight int
}

func Solution(intervals [][]int) []int {
	n := len(intervals)
	invs := make([]Interval, n)
	for i, v := range intervals {
		invs[i] = Interval{
			id:     i,
			l:      v[0],
			r:      v[1],
			weight: v[2],
		}
	}

	// 1. 按右端点升序排序；若右端点相同，按原索引升序排序（保证后续字典序处理）
	sort.Slice(invs, func(i, j int) bool {
		if invs[i].r != invs[j].r {
			return invs[i].r < invs[j].r
		}
		return invs[i].id < invs[j].id
	})

	// 预处理：对于每个区间 i，用二分查找找到在其左侧且与它不重叠的最后一个区间索引
	// 不重叠条件：r_p < l_i
	p := make([]int, n)
	for i := 0; i < n; i++ {
		left, right := 0, i-1
		idx := -1
		for left <= right {
			mid := (left + right) / 2
			if invs[mid].r < invs[i].l {
				idx = mid
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
		p[i] = idx
	}

	// dp[i][k] 表示考虑前 i 个区间，最多选 k 个不相交区间的最大权重
	// 为方便回溯，我们同时记录达到该最大权重的方案（由哪些原索引组成）
	type State struct {
		weight  int
		indices []int // 存储原索引，并保持升序
	}

	// 初始化 dp 表，大小为 (n+1) x 5
	dp := make([][]State, n+1)
	for i := range dp {
		dp[i] = make([]State, 5)
	}

	// 辅助函数：比较两个索引数组的字典序，返回较小的那个
	compareLexicographical := func(a, b []int) []int {
		if len(a) == 0 {
			return b
		}
		if len(b) == 0 {
			return a
		}
		for i := 0; i < len(a) && i < len(b); i++ {
			if a[i] < b[i] {
				return a
			} else if a[i] > b[i] {
				return b
			}
		}
		if len(a) < len(b) {
			return a
		}
		return b
	}

	// 状态转移
	for i := 1; i <= n; i++ {
		cur := invs[i-1]
		for k := 0; k <= 4; k++ {
			// 选项 1：不选当前区间 i
			bestWeight := dp[i-1][k].weight
			bestIndices := dp[i-1][k].indices

			// 选项 2：选择当前区间 i（前提是 k > 0）
			if k > 0 {
				prevIdx := p[i-1] // p 数组是 0-based，对应 invs 的下标
				prevWeight := 0
				var prevIndices []int
				if prevIdx != -1 {
					prevWeight = dp[prevIdx+1][k-1].weight
					prevIndices = dp[prevIdx+1][k-1].indices
				}

				candidateWeight := prevWeight + cur.weight
				// 组合当前的索引
				candidateIndices := make([]int, len(prevIndices)+1)
				copy(candidateIndices, prevIndices)
				// 插入并保持有序（由于是按结束时间排序加入的，直接append后排序或找位置插）
				candidateIndices[len(prevIndices)] = cur.id
				sort.Ints(candidateIndices) // 保证每次存储的索引列表是有序的

				if candidateWeight > bestWeight {
					bestWeight = candidateWeight
					bestIndices = candidateIndices
				} else if candidateWeight == bestWeight && len(candidateIndices) > 0 {
					// 权重相同时，对比字典序，选择字典序更小的索引组合
					bestIndices = compareLexicographical(bestIndices, candidateIndices)
				}
			}

			dp[i][k] = State{weight: bestWeight, indices: bestIndices}
		}
	}

	// 在所有可选的 k (0 <= k <= 4) 中找到最大权重，并保证字典序最小
	maxW := -1
	var ans []int
	for k := 0; k <= 4; k++ {
		if dp[n][k].weight > maxW {
			maxW = dp[n][k].weight
			ans = dp[n][k].indices
		} else if dp[n][k].weight == maxW && maxW != -1 {
			ans = compareLexicographical(ans, dp[n][k].indices)
		}
	}

	return ans
}
