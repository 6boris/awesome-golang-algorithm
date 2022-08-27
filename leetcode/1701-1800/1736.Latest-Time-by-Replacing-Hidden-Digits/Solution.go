package Solution

import "fmt"

func Solution(time string) string {
	questionMark := uint8(63)
	var h1, h2, m1, m2 byte
	_, _ = fmt.Sscanf(time, "%c%c:%c%c", &h1, &h2, &m1, &m2)
	if m1 == questionMark {
		m1 = '5'
	}
	if m2 == questionMark {
		m2 = '9'
	}

	if h2 == questionMark {
		// ?? -> 23
		// 0?, 1? -> 09, 19
		h2 = '3'
		if h1 == questionMark {
			h1 = '2'
		}
		if h1 < '2' {
			h2 = '9'
		}
	} else if h1 == questionMark {
		h1 = '1'
		if h2 <= '3' {
			h1 = '2'
		}
	}
	return fmt.Sprintf("%c%c:%c%c", h1, h2, m1, m2)
}
