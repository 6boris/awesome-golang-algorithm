package Solution

func Solution(arr []int) int {
	n, prev, c := len(arr), arr[0], 0
	for _, val := range arr {
		if val == prev {
			c++
			if c > n/4 {
				return prev
			}
		} else {
			prev = val
			c = 1
		}
	}
	return 0
}
