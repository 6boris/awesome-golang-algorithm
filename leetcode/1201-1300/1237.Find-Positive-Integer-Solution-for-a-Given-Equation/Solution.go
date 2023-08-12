package Solution

func Solution(customFunction func(int, int) int, z int) [][]int {
	ans := make([][]int, 0)
	for y := 1; y <= 1000; y++ {
		l, r := 1, 1000
		if r := customFunction(1, y); r > z {
			break
		}
		for l < r {
			m := l + (r-l)/2
			rr := customFunction(m, y)

			if rr == z {
				ans = append(ans, []int{m, y})
				break
			}
			if rr < z {
				l = m + 1
				continue
			}
			r = m
		}
	}
	return ans
}
