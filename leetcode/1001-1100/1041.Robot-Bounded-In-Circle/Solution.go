package Solution

var (
	turn90Degrees = []byte{'n', 'e', 's', 'w'}

	dirMap = map[byte][2]int{
		'n': {0, 1},
		's': {0, -1},
		'w': {-1, 0},
		'e': {1, 0},
	}
)

func Solution(instructions string) bool {
	x, y := 0, 0
	dirIndex := 0
	var (
		cur [2]int
		add int
	)
	for i := 0; i < 4; i++ {
		for _, i := range instructions {
			cur = dirMap[turn90Degrees[dirIndex]]
			if i == 'G' {
				x, y = x+cur[0], y+cur[1]
				continue
			}

			if i == 'L' {
				add = -1
			}
			if i == 'R' {
				add = 1
			}

			dirIndex = (dirIndex + add + 4) % 4
		}
		if x == 0 && y == 0 {
			return true
		}
	}

	return false
}
