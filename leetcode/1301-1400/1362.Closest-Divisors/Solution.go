package Solution

import "math"

func factor(n int) (int, int) {
	sn := int(math.Sqrt(float64(n)))
	a, b := 1, n
	for i := sn; i >= 1; i-- {
		if n%i != 0 {
			continue
		}
		a, b = i, n/i
		break
	}
	return a, b
}

func Solution(num int) []int {
	a1, b1 := factor(num + 1)
	a2, b2 := factor(num + 2)
	diff1 := b1 - a1
	if diff1 < 0 {
		diff1 = -diff1
	}
	diff2 := b2 - a2
	if diff2 < 0 {
		diff2 = -diff2
	}
	if diff1 < diff2 {
		return []int{a1, b1}
	}
	return []int{a2, b2}
}
