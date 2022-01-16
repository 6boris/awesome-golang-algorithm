package Solution

func Solution(s, t string) byte {
	start := uint8(0)
	for _, b := range append([]byte(s), t...) {
		start ^= b
	}
	return byte(start)
}
