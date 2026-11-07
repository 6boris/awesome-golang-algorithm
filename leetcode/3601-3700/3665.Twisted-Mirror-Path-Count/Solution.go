package Solution

func Solution(grid [][]int) int {
	const MOD = 1000000007
	m := len(grid)
	n := len(grid[0])

	// memo[r][c][dir] 记录从 (r, c) 以 dir 方向出发到达终点的路径数
	// dir: 0 表示向右 (right)，1 表示向下 (down)
	memo := make([][][]*int, m)
	for i := 0; i < m; i++ {
		memo[i] = make([][]*int, n)
		for j := 0; j < n; j++ {
			memo[i][j] = make([]*int, 2)
		}
	}

	var dfs func(r, c int, dir int) int
	dfs = func(r, c int, dir int) int {
		// 到达终点
		if r == m-1 && c == n-1 {
			return 1
		}
		if memo[r][c][dir] != nil {
			return *memo[r][c][dir]
		}

		var total int = 0

		// 模拟从当前位置和方向出发，经过可能连续的镜子反射后的最终落点
		currR, currC, currDir := r, c, dir
		for {
			// 移动一步
			if currDir == 0 { // 向右
				currC++
			} else { // 向下
				currR++
			}

			// 检查是否出界
			if currR < 0 || currR >= m || currC < 0 || currC >= n {
				total = 0
				break
			}

			// 如果到达终点
			if currR == m-1 && currC == n-1 {
				total = dfs(currR, currC, currDir) // 递归计算后续
				break
			}

			// 如果是空地，停止反射，从该空地可以继续向右或向下探索
			if grid[currR][currC] == 0 {
				// 从空地可以分叉：既可以向下走，也可以向右走
				res1 := dfs(currR, currC, 1) // 向下
				res2 := dfs(currR, currC, 0) // 向右
				total = (res1 + res2) % MOD
				break
			}

			// 如果是镜子，触发反射规则：
			// 如果原本是向右(0)进入镜子，转为向下(1)
			// 如果原本是向下(1)进入镜子，转为向右(0)
			if currDir == 0 {
				currDir = 1
			} else {
				currDir = 0
			}
			// 继续循环，模拟下一次反射
		}

		memo[r][c][dir] = &total
		return total
	}

	// 从起点 (0,0) 开始，初始可以向右(0)或向下(1)
	ansRight := dfs(0, 0, 0)
	ansDown := dfs(0, 0, 1)

	return (ansRight + ansDown) % MOD
}
