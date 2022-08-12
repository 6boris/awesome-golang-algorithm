package Solution

func Solution(p, q int) int {
	g := gcd858(p, q)
	p /= g
	q /= g

	if p&1 == 0 {
		return 2
	}
	if q&1 == 0 {
		return 0
	}

	return 1
}

func gcd858(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd858(b, a%b)
}
