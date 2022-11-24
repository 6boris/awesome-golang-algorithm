package Solution

func Solution(mat [][]int) []int {
	x, y := 0, 0
	m, n := len(mat), len(mat[0])
	idx := 0
	xDir, yDir := -1, 1
	ans := make([]int, m*n)
	for ; idx < m*n; idx++ {
		ans[idx] = mat[x][y]
		switch {
		case x == 0 && yDir == 1:
			if y < n-1 {
				y++
			} else {
				x++
			}
			xDir, yDir = 1, -1
		case y == 0 && xDir == 1:
			if x < m-1 {
				x++
			} else {
				y++
			}
			xDir, yDir = -1, 1
		case x == m-1 && yDir == -1:
			y++
			xDir, yDir = -1, 1
		case y == n-1 && xDir == -1:
			x++
			xDir, yDir = 1, -1
		default:
			x, y = x+xDir, y+yDir
		}
	}
	return ans
}
