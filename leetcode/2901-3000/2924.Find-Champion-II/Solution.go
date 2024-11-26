package Solution

func Solution(n int, edges [][]int) int {
	in := make(map[int]int)
	for _, e := range edges {
		in[e[1]]++
	}
	ans := -1
	cur := 0
	for i := range n {
		if in[i] == 0 {
			cur++
			ans = i
			if cur > 1 {
				return -1
			}
		}
	}
	return ans
}
