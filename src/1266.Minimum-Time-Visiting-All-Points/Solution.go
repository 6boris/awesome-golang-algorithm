package Solution

func Solution(points [][]int) int {
	path := 0
	for i := 0; i < len(points)-1; i++ {
		path += walk(points[i], points[i+1])
	}
	return path
}

func walk(A []int, B []int) int {
	Ax, Ay := A[0], A[1]
	Bx, By := B[0], B[1]
	return max(abs(Ax-Bx), abs(Ay-By))
}

func abs(val int) int {
	if val < 0 {
		return -val
	}
	return val
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
