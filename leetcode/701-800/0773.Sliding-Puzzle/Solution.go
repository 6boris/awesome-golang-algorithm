package Solution

func isOk773(board [2][3]int) bool {
	return board[0][0] == 1 && board[0][1] == 2 && board[0][2] == 3 &&
		board[1][0] == 4 && board[1][1] == 5
}

type qItem773 struct {
	key       [2][3]int
	zeroIndex [2]int
}

var dir773 = [4][2]int{
	{0, 1}, {0, -1}, {1, 0}, {-1, 0},
}

func Solution(board [][]int) int {

	key := [2][3]int{
		{board[0][0], board[0][1], board[0][2]},
		{board[1][0], board[1][1], board[1][2]},
	}
	zeroIndex := [2]int{0, 0}
	for i := range 2 {
		for j := range 3 {
			if board[i][j] == 0 {
				zeroIndex = [2]int{i, j}
			}
		}
	}
	queue := []qItem773{
		{key, zeroIndex},
	}
	used := map[[2][3]int]struct{}{
		key: struct{}{},
	}
	steps := 0
	for len(queue) > 0 {
		nq := make([]qItem773, 0)
		for _, cur := range queue {
			if isOk773(cur.key) {
				return steps
			}
			x, y := cur.zeroIndex[0], cur.zeroIndex[1]
			for _, d := range dir773 {
				nx, ny := x+d[0], y+d[1]
				if nx >= 0 && nx <= 1 && ny >= 0 && ny <= 2 {
					b := cur.key
					b[nx][ny], b[x][y] = b[x][y], b[nx][ny]
					if _, ok := used[b]; !ok {
						nq = append(nq, qItem773{b, [2]int{nx, ny}})
						used[b] = struct{}{}
					}
				}
			}
		}
		queue = nq
		steps++
	}
	return -1
}
