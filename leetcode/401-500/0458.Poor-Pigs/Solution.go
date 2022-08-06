package Solution

func Solution(buckets int, minutesToDie int, minutesToTest int) int {
	base := minutesToTest / minutesToDie
	return powerX(base+1, buckets)
}

func powerX(a, n int) int {
	p := 0
	base := 1
	for ; base < n; base *= a {
		p++
	}
	return p
}
