package Solution

const mod1301 = 1000000007

type pathItem struct {
	score, cnt int
}

func Solution(board []string) []int {
	m, n := len(board), len(board[0])
	cache := make([][]pathItem, m)
	for i := range n {
		cache[i] = make([]pathItem, n)
		for j := range n {
			cache[i][j] = pathItem{-1, 0}
		}
	}

	cache[m-1][n-1] = pathItem{0, 1}
	cache[0][0] = pathItem{0, 0}

	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			if board[i][j] == 'X' || board[i][j] == 'S' {
				continue
			}
			right := j + 1
			down := i + 1
			score := 0
			if i != 0 || j != 0 {
				score = int(board[i][j] - '0')
			}
			if right < n && cache[i][right].score != -1 {
				try := cache[i][right].score + score
				if try > cache[i][j].score {
					cache[i][j].score = try
					cache[i][j].cnt = cache[i][right].cnt
				} else if try == cache[i][j].score {
					cache[i][j].cnt = (cache[i][j].cnt + cache[i][right].cnt) % mod1301
				}
			}

			if down < m && cache[down][j].score != -1 {
				try := cache[down][j].score + score
				if try > cache[i][j].score {
					cache[i][j].score = try
					cache[i][j].cnt = cache[down][j].cnt
				} else if try == cache[i][j].score {
					cache[i][j].cnt = (cache[i][j].cnt + cache[down][j].cnt) % mod1301
				}
			}

			if down < m && right < n && cache[down][right].score != -1 {
				try := cache[down][right].score + score
				if try > cache[i][j].score {
					cache[i][j].score = try
					cache[i][j].cnt = cache[down][right].cnt
				} else if try == cache[i][j].score {
					cache[i][j].cnt = (cache[i][j].cnt + cache[down][right].cnt) % mod1301
				}
			}
		}
	}
	return []int{cache[0][0].score, cache[0][0].cnt}
}
