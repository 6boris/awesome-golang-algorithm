package Solution

import "math"

func divide(dividend int, divisor int) int {
	dm, dsm, sign := true, true, true
	if dividend < 0 {
		dm = false
		dividend = -dividend
	}
	if divisor < 0 {
		dsm = false
		divisor = -divisor
	}

	if (dm && !dsm) || (!dm && dsm) {
		sign = false
	}

	if divisor == 0 {
		return math.MaxInt32
	}

	h := 0
	out := 0
	if dividend < divisor {
		return 0
	}
	for (divisor << uint(h)) <= dividend {
		h++
	}
	h--
	for i := h; i >= 0; i-- {
		if dividend >= (divisor << uint(i)) {
			dividend -= (divisor << uint(i))
			out |= 1 << uint(i)
		}
	}

	if !sign {
		out = -out
	}
	if out > math.MaxInt32 {
		return math.MaxInt32
	} else if out < math.MinInt32 {
		return math.MinInt32
	} else {
		return out
	}
}
