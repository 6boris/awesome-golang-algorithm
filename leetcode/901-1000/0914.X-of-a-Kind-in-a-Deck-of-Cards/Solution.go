package Solution

func Solution(deck []int) bool {
	n := len(deck)
	if n < 2 {
		return false
	}
	hash := make(map[int]int)
	for _, val := range deck {
		hash[val]++
	}
	ans := hash[deck[0]]
	for _, val := range hash {
		ans = gcd(ans, val)
	}
	return ans >= 2
}

func gcd(x, y int) int {
	for y > 0 {
		x, y = y, x%y
	}
	return x
}
