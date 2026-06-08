package Solution

func Solution(a, b, c int) int {
	total := a + b + c
	max_heap := max(a, b, c)
	return min(total/2, total-max_heap)
}
