package Solution

func Solution(s string) int {
	a := make([]int, len(s))
	at := 0
	ans := 0
	// () ( () )
	// 1 0 1
	for i := 1; i < len(s); i++ {
		if s[i] == '(' {
			at++
			a[at] = 0
			continue
		}
		if a[at] == 0 {
			a[at] = 1
			continue
		}
		tmp := 0
		for ; at != 0 && a[at] != 0; at-- {
			tmp += a[at]
		}
		a[at] = tmp * 2
	}
	for i := 0; i <= at; i++ {
		ans += a[i]
	}
	return ans
}
