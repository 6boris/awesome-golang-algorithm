package Solution

import "fmt"

func Solution(low int, high int) int {
	var ok func(int) bool
	ok = func(n int) bool {
		s := fmt.Sprintf("%d", n)
		l := len(s)
		if l&1 == 1 {
			return false
		}
		// 1 2, 3, 4
		left, right := 0, 0
		for start, end := 0, l-1; start < end; start, end = start+1, end-1 {
			left += int(s[start] - '0')
			right += int(s[end] - '0')
		}
		return left == right
	}
	cnt := 0
	for i := low; i <= high; i++ {
		if ok(i) {
			cnt++
		}
	}
	return cnt
}
