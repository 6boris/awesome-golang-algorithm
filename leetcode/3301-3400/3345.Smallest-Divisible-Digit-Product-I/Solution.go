package Solution

func Solution(n int, t int) int {
	p, c := 1, n
	for ; c > 0; c /= 10 {
		p *= c % 10
	}
	if p%t == 0 {
		return n
	}

	digit := n % 10
	p /= digit

	for i := digit + 1; i <= 9; i++ {
		x := p * i
		if x%t == 0 {
			return n + i - digit
		}
	}
	return n + 10 - digit
}
