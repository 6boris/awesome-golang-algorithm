package Solution

func Solution(moves string) bool {
	x, y := 0, 0

	for _, b := range []byte(moves) {
		switch b {
		case 'U':
			x--
		case 'D':
			x++
		case 'R':
			y++
		case 'L':
			y--
		default:
			return false
		}
	}
	return x == 0 && y == 0
}
