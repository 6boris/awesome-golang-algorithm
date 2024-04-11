package Solution

import "fmt"

func Solution(n int) int {
	ns := []byte(fmt.Sprintf("%d", n))
	idx := 1
	for ; idx < len(ns); idx++ {
		if ns[idx] < ns[idx-1] {
			break
		}
	}
	if idx != len(ns) {
		index := idx - 1
		top := ns[index]
		pre := index - 1
		for ; pre >= 0 && ns[pre] == top; pre-- {
			ns[pre] = '9'
		}
		ns[index] = '9'
		ns[pre+1] = top - 1
		for ; idx < len(ns); idx++ {
			ns[idx] = '9'
		}
	}

	ans := 0
	for i := 0; i < len(ns); i++ {
		ans = ans*10 + int(ns[i]-'0')
	}
	return ans
}
