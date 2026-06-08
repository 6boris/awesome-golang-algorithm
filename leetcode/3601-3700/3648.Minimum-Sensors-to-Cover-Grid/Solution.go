package Solution

func Solution(n, m, k int) int {
	x := 2*k + 1
	rows := (n + x - 1) / x
	cols := (m + x - 1) / x
	return rows * cols
}
