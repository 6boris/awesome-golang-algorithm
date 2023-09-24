package Solution

func Solution(land [][]int) [][]int {
	ans := make([][]int, 0)
	rows, cols := len(land), len(land[0])
	var findEnd func(int, int) (int, int)
	findEnd = func(x, y int) (int, int) {

		x1, y1 := x, y
		for ; y1 < cols && land[x][y1] != 0; y1++ {
		}
		y1--
		for ; x1 < rows && land[x1][y] != 0; x1++ {
			for l := y; l <= y1; l++ {
				land[x1][l] = 0
			}
		}
		x1--
		return x1, y1
	}
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if land[r][c] == 0 {
				continue
			}
			s, e := findEnd(r, c)
			ans = append(ans, []int{r, c, s, e})
		}
	}
	return ans
}
