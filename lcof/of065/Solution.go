package Solution

func add(a int, b int) int {
	for b != 0 {
		carry := (a & b) << 1
		n := a ^ b
		a = n
		b = carry
	}
	return a
}
