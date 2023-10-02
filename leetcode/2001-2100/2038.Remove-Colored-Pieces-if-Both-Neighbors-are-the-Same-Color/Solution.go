package Solution

func Solution(colors string) bool {
	a, b := 0, 0
	x := 0
	for _, b := range []byte(colors) {
		if b == 'B' {
			if x >= 3 {
				a += x - 2
			}
			x = 0
			continue
		}
		x++
	}
	if x >= 3 {
		a += x - 2
	}
	x = 0
	for _, bb := range []byte(colors) {
		if bb == 'A' {
			if x >= 3 {
				b += x - 2
			}
			x = 0
			continue
		}
		x++
	}
	if x >= 3 {
		b += x - 2
	}
	return a > b
}
