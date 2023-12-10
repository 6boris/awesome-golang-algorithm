package Solution

func Solution(target int) int {
	if target < 0 {
		target = -target
	}
	n := 0
	for target > 0 {
		n++
		target -= n
	}
	if target&1 == 0 {
		return n
	}
	return n + 1 + n&1
}
