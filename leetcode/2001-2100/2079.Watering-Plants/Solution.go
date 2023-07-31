package Solution

func Solution(plants []int, capacity int) int {
	cost := 0
	sum := 0
	for idx, water := range plants {
		sum += water
		if sum <= capacity {
			cost++
			continue
		}
		cost += 2*idx + 1
		sum = water
	}
	return cost
}
