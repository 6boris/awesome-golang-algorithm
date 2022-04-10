package Solution

import "strconv"

func Solution(ops []string) int {
	ans := 0

	points := make([]int, len(ops))
	idx := 0
	for i := 0; i < len(ops); i, idx = i+1, idx+1 {
		op := ops[i]
		if op == "C" {
			idx -= 2
			continue
		}
		if op == "D" {
			points[idx] = points[idx-1] * 2
			continue
		}

		if op == "+" {
			points[idx] = points[idx-1] + points[idx-2]
			continue
		}
		_t, _ := strconv.Atoi(op)
		points[idx] = _t
	}
	for i := 0; i < idx; i++ {
		ans += points[i]
	}
	return ans
}
