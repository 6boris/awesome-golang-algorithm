package Solution

func Solution(n int) bool {
	sum, prod := 0, 1
	var tmp int
	for base := n; base > 0; base /= 10 {
		tmp = base % 10
		sum += tmp
		prod *= tmp
	}
	return n%(sum+prod) == 0
}
