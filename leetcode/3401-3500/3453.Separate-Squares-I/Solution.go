package Solution

import (
	"math"
)

func Solution(squares [][]int) float64 {
	max_y, total_area := 0.0, 0.0
	for _, sq := range squares {
		y, l := sq[1], sq[2]
		total_area += float64(l * l)
		if float64(y+l) > max_y {
			max_y = float64(y + l)
		}
	}

	check := func(limit_y float64) bool {
		area := 0.0
		for _, sq := range squares {
			y, l := sq[1], sq[2]
			if float64(y) < limit_y {
				overlap := math.Min(limit_y-float64(y), float64(l))
				area += float64(l) * overlap
			}
		}

		return area >= total_area/2.0
	}

	lo, hi := 0.0, max_y
	eps := 1e-5
	for math.Abs(hi-lo) > eps {
		mid := (hi + lo) / 2.0
		if check(mid) {
			hi = mid
		} else {
			lo = mid
		}
	}
	return hi
}
