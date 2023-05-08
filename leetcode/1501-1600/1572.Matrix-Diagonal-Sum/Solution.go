package Solution

func Solution(mat [][]int) int {
	length := len(mat)
	x, y := 0, 0
	x1, y1 := 0, length-1
	ans := 0
	for x < length {
		ans += mat[x][y]
		if x1 != x || y1 != y {
			ans += mat[x1][y1]
		}
		x, y = x+1, y+1
		x1, y1 = x1+1, y1-1
	}
	return ans
}
