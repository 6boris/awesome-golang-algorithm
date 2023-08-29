package Solution

func Solution(customers string) int {
	l := len(customers)
	left := make([]int, l+1)
	right := make([]int, l+1)
	n := 0
	for closeTime := 0; closeTime < l; closeTime++ {
		if customers[closeTime] == 'Y' {
			left[closeTime] = n + 1
			continue
		}
		left[closeTime] = n
		n++
	}
	left[l] = n
	y := 0
	for closeTime := l - 1; closeTime >= 0; closeTime-- {
		right[closeTime] = y
		if customers[closeTime] == 'Y' {
			y++
		}
	}
	ans := -1
	for t := 0; t <= l; t++ {
		if ans == -1 || left[t]+right[t] < left[ans]+right[ans] {
			ans = t
		}
	}
	return ans
}
