package Solution

func Solution(croakOfFrogs string) int {
	if croakOfFrogs[0] != 'c' {
		return -1
	}
	indies := map[byte]int{
		'c': 0, 'r': 1, 'o': 2, 'a': 3, 'k': 4,
	}
	ret, free := 0, 0
	state := [5]int{}
	for _, b := range []byte(croakOfFrogs) {
		// croak
		switch b {
		case 'c':
			// 一个新青蛙
			if free > 0 {
				free--
			} else {
				ret++
			}
			state[indies[b]]++
		case 'r', 'o', 'a':
			if state[indies[b]-1] == 0 {
				return -1
			}
			state[indies[b]-1]--
			state[indies[b]]++
		case 'k':
			if state[indies[b]-1] == 0 {
				return -1
			}
			state[indies[b]-1]--
			free++
		default:
			return -1
		}
	}
	for i := 0; i < 5; i++ {
		if state[i] > 0 {
			return -1
		}
	}

	return ret
}
