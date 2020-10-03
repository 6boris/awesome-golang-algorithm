package Solution

func Solution(bits []int) bool {
	n := len(bits)
	if n == 1 {
		return true
	}
	for i := 0; i < n; {
		if i == n-1 {
			return true
		}
		if bits[i] == 0 {
			i++
		} else {
			i += 2
		}
	}
	return false
}
