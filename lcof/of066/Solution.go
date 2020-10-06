package Solution

func constructArr(a []int) []int {
	left, right, ans := 1, 1, make([]int, len(a))
	for i := 0; i < len(a); i++ {
		ans[i] = left
		left *= a[i]
	}
	for i := len(a) - 1; i >= 0; i-- {
		ans[i] *= right
		right *= a[i]
	}
	return ans
}
