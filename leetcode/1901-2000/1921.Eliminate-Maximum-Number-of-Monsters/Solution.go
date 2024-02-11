package Solution

import (
	"sort"
)

func Solution(dist []int, speed []int) int {
	minutes := make([]float64, len(dist))
	for i := 0; i < len(dist); i++ {
		minutes[i] = float64(dist[i]) / float64(speed[i])
	}
	sort.Float64s(minutes)

	cost := 0
	ans := 0
	full := true
	for i := 0; i < len(minutes); i++ {
		if full {
			ans++
			full = false
			continue
		}
		// need charge
		leftMinutes := minutes[i] - float64(cost)
		if leftMinutes <= 1.0 {
			break
		}
		i--
		full = true
		cost++
	}
	return ans
}
