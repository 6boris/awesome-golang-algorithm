package Solution

func Solution(m int, n int, ops [][]int) int {
	minX, minY := m, n
	for _, op := range ops {
		minX = min(minX, op[0])
		minY = min(minY, op[1])
	}
	return minX * minY
}
