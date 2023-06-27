package Solution

func Solution(garbage []string, travel []int) int {
	g, m, p := 0, 0, 0
	cost := 0
	for i := len(garbage) - 1; i >= 0; i-- {
		for _, _g := range garbage[i] {
			if _g == 'G' {
				g++
			}
			if _g == 'M' {
				m++
			}
			if _g == 'P' {
				p++
			}
		}
		if g > 0 && i != 0 {
			cost += travel[i-1]
		}
		if m > 0 && i != 0 {
			cost += travel[i-1]
		}
		if p > 0 && i != 0 {
			cost += travel[i-1]
		}
	}
	return g + m + p + cost
}
