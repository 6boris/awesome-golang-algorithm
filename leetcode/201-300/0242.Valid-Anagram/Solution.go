package Solution

// 用异或计算
func isAnagram_3(s, t string) bool {
	if len(s) != len(t) {
		return false
	}
	var x, y rune

	for _, v := range s {
		x ^= v
		y += v * v
	}
	for _, v := range t {
		x ^= v
		y -= v * v
	}
	return x == 0 && y == 0
}
