package Solution

func Solution(S string) [][]int {
	n := len(S)
	if n < 3 {
		return [][]int{}
	}
	res := make([][]int, 0)
	c, l, r := 1, -1, -1
	for i := 1; i < n; i++ {
		if S[i] == S[i-1] {
			c++
			if l == -1 {
				l = i - 1
			}
			if c >= 3 {
				r = i
			}
		} else {
			if c >= 3 {
				res = append(res, []int{l, r})
			}
			c, l = 1, -1
		}
	}
	if c >= 3 {
		res = append(res, []int{l, r})
	}
	return res
}
