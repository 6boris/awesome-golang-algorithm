package Solution

/*
0
0          1
0    1     1     0
0 1  1  0  1  0  0  1
1 2  3  4  5  6  7  8
*/
func Solution(n int, k int) int {
	nn := n
	path := make([]int, 0)
	for nn > 0 {
		path = append(path, k)
		if k&1 == 1 {
			k++
		}
		k /= 2
		nn--
	}
	ans := 0
	for i := len(path) - 2; i >= 0; i-- {
		a, b := 0, 1
		if ans != 0 {
			a, b = 1, 0
		}
		if path[i] == path[i+1]*2 {
			ans = b
			continue
		}
		ans = a
	}
	return ans
}
