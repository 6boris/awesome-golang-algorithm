package Solution

import "fmt"

func Solution(arr []int) string {
	in := make(map[int]int)
	for _, n := range arr {
		in[n]++
	}
	var a, b, c, d int
	for i := 23; i >= 0; i-- {
		a, b = i/10, i%10
		if a == b {
			if in[a] < 2 {
				continue
			}
		} else if in[a] < 1 || in[b] < 1 {
			continue
		}
		in[a]--
		in[b]--

		for j := 59; j >= 0; j-- {
			c, d = j/10, j%10
			if c == d {
				if in[c] >= 2 {
					return fmt.Sprintf("%d%d:%d%d", a, b, c, d)
				}
			} else if in[c] > 0 && in[d] > 0 {
				return fmt.Sprintf("%d%d:%d%d", a, b, c, d)
			}
		}
		in[a]++
		in[b]++
	}
	return ""
}
