package Solution

func Solution(grid []string) int {
	// 感觉类似灌水的思想，设置一个nxn的数组，开始数字都是0，
	// 从没有被访问的格子开始dfs，大概就可以知道有多少个联通的去了
	// 标记+dsf
	/*
	  对于一个1x1的放个，如果如果放置的是 / ,左上角标记为1，右下角标记4
	  如果放置\, 右上角是2，左下角是3，
	  所以如果一个放个完全被用过后，他的值应该是5
	*/
	n := len(grid)
	m := make([][]int, n)
	ch := make([][]byte, n)

	for i := 0; i < n; i++ {
		ch[i] = make([]byte, n)
		m[i] = make([]int, n)
	}

	for r, line := range grid {
		for c, b := range line {
			ch[r][c] = byte(b)
		}
	}
	// 1v4, 2v3
	var dfs func(int, int, int)
	// 只需要根据数字就可以知道如何走了。
	dfs = func(x, y, dir int) {
		// 最基础的判断方式
		if x < 0 || x >= n || y < 0 || y >= n || m[x][y] == 5 {
			return
		}
		//例如如果dfs是否从下面来的，就需要判断当前网格的值是是否是3，或者4，如果是表示便利过
		// 其他方向同理，

		if m[x][y] == 0 {
			// 还没有访问过
			if ch[x][y] == ' ' {
				// 空格直接灌满
				m[x][y] = 5
				dfs(x, y-1, 3)
				dfs(x, y+1, 1)
				dfs(x-1, y, 4)
				dfs(x+1, y, 2)
				return
			}
			if ch[x][y] == '/' {
				// 标记为1
				if dir == 1 || dir == 2 {
					m[x][y] = 1
					dfs(x-1, y, 4)
					//  坑就在这里
					dfs(x, y-1, 3)
				} else {
					m[x][y] = 4
					dfs(x, y+1, 1)
					dfs(x+1, y, 2)
				}
			} else {
				if dir == 2 || dir == 3 {
					m[x][y] = 2
					dfs(x-1, y, 4)
					dfs(x, y+1, 1)

				} else {
					m[x][y] = 3
					dfs(x, y-1, 3)
					dfs(x+1, y, 2)
				}
			}
			return
		}
		// dir=1表示从左侧来，2表示上方，3表示右侧，4表示下方
		if m[x][y] == 1 {
			// 遍历过了
			if dir == 1 || dir == 2 {
				return
			}
			// 左上角，因为还不等于5， 只能走右,下
			m[x][y] = 5
			dfs(x, y+1, 1)
			dfs(x+1, y, 2)
			return
		}
		if m[x][y] == 4 {
			if dir == 3 || dir == 4 {
				return
			}
			m[x][y] = 5
			// 右上角，只能走左，下
			dfs(x, y-1, 3)
			dfs(x-1, y, 4)
			return

		}
		if m[x][y] == 2 {
			if dir == 2 || dir == 3 {
				return
			}
			m[x][y] = 5
			dfs(x, y-1, 3)
			dfs(x+1, y, 2)
			return

		}
		if m[x][y] == 3 {
			if dir == 1 || dir == 4 {
				return
			}
			m[x][y] = 5
			dfs(x-1, y, 4)
			dfs(x, y+1, 1)
			return
		}
	}

	ans := 0
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if m[i][j] != 5 {
				dfs(i, j, 0)
				ans++
			}
			if m[i][j] != 5 {
				dfs(i, j, 0)
				ans++
			}
		}
	}

	//  开始哦都市0，占用了一半

	return ans
}
