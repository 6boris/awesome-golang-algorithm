package Solution

func Solution(left, right int) []int {
	var res []int
	var ok func(int) bool = func(n int) bool {
		self := n
		for n > 0 {
			mod := n % 10
			if mod == 0 {
				return false
			}
			if self%mod != 0 {
				return false
			}
			n /= 10
		}
		return true
	}
	for i := left; i <= right; i++ {
		if ok(i) {
			res = append(res, i)
		}
	}
	return res
}
