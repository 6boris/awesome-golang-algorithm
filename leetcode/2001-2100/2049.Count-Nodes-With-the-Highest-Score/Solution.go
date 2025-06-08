package Solution

func Solution(parents []int) int {
	adj := make(map[int][]int)
	for c, f := range parents {
		adj[f] = append(adj[f], c)
	}
	n := len(parents)
	children := make([][2]int, n)
	var dfs func(int) int
	dfs = func(root int) int {
		children[root] = [2]int{-1, -1}
		index := 0
		count := 0
		for _, child := range adj[root] {
			children[root][index] = dfs(child)
			count += children[root][index]
			index++
		}
		return count + 1
	}
	_ = dfs(0)
	ans := 0
	// 尝试移除0
	left := children[0][0]
	if left == -1 {
		left = 1
	}
	right := children[0][1]
	if right == -1 {
		right = 1
	}
	ans = left * right
	count := 1
	for i := 1; i < n; i++ {
		x := n
		x--
		left = children[i][0]
		right = children[i][1]
		if left != -1 {
			x -= left
		} else {
			left = 1
		}
		if right != -1 {
			x -= right
		} else {
			right = 1
		}
		now := left * right * x
		if now == ans {
			count++
		} else if now > ans {
			ans = now
			count = 1
		}
	}
	return count
}
