package Solution

const (
	rightMask = 0x1E
	leftMask  = 0x1E0
	midMask   = 0x78
)

func Solution(n int, reservedSeats [][]int) int {
	// 10^4个操作
	bited := make(map[int]int)
	for _, seat := range reservedSeats {
		bited[seat[0]] |= 1 << (seat[1] - 1)
	}
	ret := (n - len(bited)) * 2

	var left, right, mid bool
	for _, seated := range bited {
		left = seated&leftMask == 0
		right = seated&rightMask == 0
		if left && right {
			ret += 2
			continue
		}
		mid = seated&midMask == 0
		if mid || left || right {
			ret++
		}
	}

	return ret
}
