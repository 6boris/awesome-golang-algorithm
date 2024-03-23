package Solution

func Solution(n int, presses int) int {
	// 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 1, 1, 2
	// 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1
	// 0,    2,    2,    2,    2,    2,    2
	// 0, 3,    3,    3,    3,    3     3,
	// 0, 4,       4,       4,       4
	n %= 12
	if n == 0 {
		n = 12
	}
	base := uint16((1 << n) - 1)

	// 0 0 1 1 1 1 1 1
	// 0 0 6 5 4 3 2 1
	r := make(map[uint16]struct{})
	basicMask := uint16(1<<n) - 1
	// 0b0001111111,  0b1010101010101010,  0b0101010101010101, 0b1001001001001001
	magicNumber := []uint16{basicMask, 0xAAAA & basicMask, 0x5555 & basicMask, 0x9249 & basicMask}

	cache := make(map[[2]uint16]struct{})
	var dfs func(uint16, uint16)
	dfs = func(cur, left uint16) {
		if left == 0 {
			r[cur] = struct{}{}
			return
		}
		if _, ok := cache[[2]uint16{cur, left}]; ok {
			return
		}
		for _, op := range magicNumber {
			dfs(cur^op, left-1)
		}
		cache[[2]uint16{cur, left}] = struct{}{}
	}
	dfs(base, uint16(presses))
	return len(r)
}
