package Solution

func Solution(n int, startPos []int, s string) []int {
	ans := make([]int, len(s))
	var steps func(int, int, int) int

	steps = func(x, y, index int) int {
		ans := 0
		for i := index; i < len(s); i++ {
			a, b := 0, 0
			if s[i] == 'L' {
				a, b = 0, -1
			} else if s[i] == 'R' {
				a, b = 0, 1
			} else if s[i] == 'U' {
				a, b = -1, 0
			} else {
				a, b = 1, 0
			}
			nx, ny := x+a, y+b
			if nx < 0 || nx >= n || ny < 0 || ny >= n {
				return ans
			}
			x, y = nx, ny
			ans++
		}
		return ans
	}
	for i := 0; i < len(s); i++ {
		ans[i] = steps(startPos[0], startPos[1], i)
	}
	return ans
}
