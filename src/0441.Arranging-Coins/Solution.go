package Solution

import "math"

func arrangeCoins(n int) int {
	if n <= 0 {
		return 0
	}

	x := math.Sqrt(2*float64(n)+0.25) - 0.5
	return int(x)
}
