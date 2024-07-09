package Solution

func Solution(customers [][]int) float64 {
	sum, end := 0, 0
	for _, c := range customers {
		sum += c[1]
		if c[0] < end {
			sum += end - c[0]
		}

		end = max(end+c[1], c[0]+c[1])
	}
	return float64(sum) / float64(len(customers))
}
