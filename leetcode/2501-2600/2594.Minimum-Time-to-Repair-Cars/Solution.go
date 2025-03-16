package Solution

import (
	"math"
	"slices"
)

func Solution(ranks []int, cars int) int64 {
	var ok func(int64) bool
	ok = func(minutes int64) bool {
		c := int64(0)
		for _, r := range ranks {
			nn := minutes / int64(r)
			i := int64(math.Sqrt(float64(nn)))
			c += i
		}
		return c >= int64(cars)
	}

	m := slices.Max(ranks)
	l, r := int64(0), int64(m)*int64(cars)*int64(cars)+1
	for l < r {
		m := (l + r) / 2
		if ok(m) {
			r = m
			continue
		}
		l = m + 1
	}
	return l
}
