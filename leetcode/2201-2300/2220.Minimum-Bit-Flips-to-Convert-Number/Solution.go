package Solution

func Solution(start, goal int) int {
	x := start ^ goal
	n := 0
	for x > 0 {
		n++
		x = x & (x - 1)
	}
	return n
}
