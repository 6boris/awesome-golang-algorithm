package Solution

func Solution(edges [][]int) int {
	a, b := edges[0], edges[1]
	if a[0] == b[1] || a[0] == b[0] {
		return a[0]
	}
	if a[1] == b[1] || a[1] == b[0] {
		return a[1]
	}
	return -1
}
