package Solution

func Solution(num int) int {
	bits := [4]int{}
	idx, six := 0, -1
	source := num

	for num > 0 {
		n := num % 10
		bits[idx] = n
		if n == 6 {
			six = idx
		}
		idx, num = idx+1, num/10
	}
	if six == -1 {
		return source
	}
	base := 0
	bits[six] = 9
	for i := idx - 1; i >= 0; i-- {
		base = base*10 + bits[i]
	}
	return base
}
