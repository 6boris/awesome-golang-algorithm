package Solution

func Solution(a, b int) int {
	for b != 0 {
		carry := (a & b) << 1
		a = a ^ b
		b = carry
	}

	return a
}
