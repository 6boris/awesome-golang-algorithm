package Solution

import "sort"

type Pair354 struct {
	val int
	idx int
}

func Solution(n int, nums []int, maxDiff int, queries [][]int) []int {
	// 1. 带着原始下标排序
	pairs := make([]Pair354, n)
	for i := 0; i < n; i++ {
		pairs[i] = Pair354{val: nums[i], idx: i}
	}
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].val < pairs[j].val
	})

	// 记录原始下标 -> 排序后的新下标
	pos := make([]int, n)
	for i := 0; i < n; i++ {
		pos[pairs[i].idx] = i
	}

	// 2. 并查集：用于 O(1) 判断两点是否在同一个大连通块内
	parent := make([]int, n)
	for i := 0; i < n; i++ {
		parent[i] = i
	}
	var find func(int) int
	find = func(i int) int {
		if parent[i] == i {
			return i
		}
		parent[i] = find(parent[i])
		return parent[i]
	}
	union := func(i, j int) {
		rootI, rootJ := find(i), find(j)
		if rootI != rootJ {
			parent[rootI] = rootJ
		}
	}

	// 3. 构建倍增表
	// up[i][j] 表示从排序下标 i 出发，向右跳 2^j 步到达的排序下标
	const maxLog = 18 // 2^17 = 131072 > 10^5
	up := make([][]int, n)
	for i := 0; i < n; i++ {
		up[i] = make([]int, maxLog)
	}

	// 用双指针初始化 up[i][0] (即跳 1 步，向右能到的最远位置)
	r := 0
	for l := 0; l < n; l++ {
		for r < n && pairs[r].val-pairs[l].val <= maxDiff {
			r++
		}
		up[l][0] = r - 1 // 满足条件的最右边界
		if l > 0 && pairs[l].val-pairs[l-1].val <= maxDiff {
			union(l, l-1)
		}
	}

	// 填充倍增表其余部分
	for j := 1; j < maxLog; j++ {
		for i := 0; i < n; i++ {
			up[i][j] = up[up[i][j-1]][j-1]
		}
	}

	// 4. 回答查询
	ans := make([]int, len(queries))
	for k, q := range queries {
		u, v := pos[q[0]], pos[q[1]]
		if u == v {
			ans[k] = 0
			continue
		}
		// 确保 u 在 v 的左边
		if u > v {
			u, v = v, u
		}

		// 检查连通性，不在一个连通块直接不可达
		if find(u) != find(v) {
			ans[k] = -1
			continue
		}

		// 如果本身一步就能跨过去
		if pairs[v].val-pairs[u].val <= maxDiff {
			ans[k] = 1
			continue
		}

		// 利用倍增向上跳跃，寻找最少步数
		steps := 0
		curr := u
		for j := maxLog - 1; j >= 0; j-- {
			// 如果跳 2^j 步之后依然到不了 v 的左边邻域（即跳完之后最远依然够不着 v）
			if up[curr][j] < v && pairs[v].val-pairs[up[curr][j]].val > maxDiff {
				steps += (1 << j)
				curr = up[curr][j]
			}
		}

		// 此时 curr 是跳跃能到达的、且距离 v 刚好还差“最后一步”的最远点
		// 再往前迈出一步即可到达能覆盖 v 的点
		steps++ // 迈出那一步
		curr = up[curr][0]

		if curr >= v || pairs[v].val-pairs[curr].val <= maxDiff {
			ans[k] = steps + 1 // 加上最后直接连到 v 的那 1 步
		} else {
			ans[k] = -1
		}
	}

	return ans
}
