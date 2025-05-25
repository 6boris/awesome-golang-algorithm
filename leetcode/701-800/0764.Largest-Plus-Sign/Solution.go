package Solution

type l764 struct {
	up, down, left, right int
}

func Solution(n int, mines [][]int) int {
	var ans int
	checker := make(map[int]map[int]struct{})
	for _, m := range mines {
		if _, ok := checker[m[0]]; !ok {
			checker[m[0]] = map[int]struct{}{}
		}
		checker[m[0]][m[1]] = struct{}{}
	}
	order := make([][]l764, n)
	for i := range n {
		order[i] = make([]l764, n)
	}

	for i := 0; i < n; i++ {
		one := 0
		for j := 0; j < n; j++ {
			if v, ok := checker[i]; ok {
				if _, ok1 := v[j]; ok1 {
					one = 0
					continue
				}
			}
			order[i][j].left = one
			one++
		}
		one = 0
		for j := n - 1; j >= 0; j-- {
			if v, ok := checker[i]; ok {
				if _, ok1 := v[j]; ok1 {
					one = 0
					continue
				}
			}
			order[i][j].right = one
			one++

		}

		one = 0
		for j := 0; j < n; j++ {
			if v, ok := checker[j]; ok {
				if _, ok1 := v[i]; ok1 {
					one = 0
					continue
				}
			}
			order[j][i].up = one
			one++
		}
		one = 0
		for j := n - 1; j >= 0; j-- {
			if v, ok := checker[j]; ok {
				if _, ok1 := v[i]; ok1 {
					one = 0
					continue
				}
			}
			order[j][i].down = one
			one++
		}
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if v, ok := checker[i]; ok {
				if _, ok1 := v[j]; ok1 {
					continue
				}
			}

			ans = max(ans, min(order[i][j].down, order[i][j].up, order[i][j].left, order[i][j].right)+1)
		}
	}
	return ans
}
