package Solution

func Solution(A []int) []bool {
	n := len(A)
	if n == 1 {
		if A[0] == 0 {
			return []bool{true}
		} else {
			return []bool{false}
		}
	}
	ans, v := make([]bool, n), 0
	for i := 0; i < n; i++ {
		v = (2*v + A[i]) % 5
		if v == 0 {
			ans[i] = true
		} else {
			ans[i] = false
		}
	}
	return ans
}
