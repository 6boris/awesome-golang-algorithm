package Solution

func Solution(dimensions [][]int) int {
	diagonal, area := 0, 0
	for _, n := range dimensions {
		tmp := n[0]*n[0] + n[1]*n[1]
		if tmp == diagonal {
			area = max(area, n[0]*n[1])
		}
		if tmp > diagonal {
			diagonal = tmp
			area = n[0] * n[1]
		}
	}

	return area
}
