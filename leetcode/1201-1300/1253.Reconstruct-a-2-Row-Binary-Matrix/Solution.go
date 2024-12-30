package Solution

func Solution(upper int, lower int, colsum []int) [][]int {
	res := make([][]int, 2)
	l := len(colsum)
	sum := 0
	for i := range l {
		sum += colsum[i]
	}
	if sum != upper+lower {
		return nil
	}

	for i := range 2 {
		res[i] = make([]int, l)
	}

	for i := range l {
		if colsum[i] == 2 {
			if upper == 0 || lower == 0 {
				return nil
			}
			res[0][i] = 1
			res[1][i] = 1
			upper--
			lower--
		}
	}

	for i := range l {
		if colsum[i] == 1 {
			if upper > 0 {
				res[0][i] = 1
				upper--
				continue
			}
			if lower > 0 {
				res[1][i] = 1
				lower--
				continue
			}
			return nil
		}
	}

	return res
}
