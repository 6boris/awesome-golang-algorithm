package Solution

func Solution(A []int) bool {
	n := len(A)
	if n <= 2 {
		return false
	}
	pass := 0
	for ; pass+1 < n && A[pass] < A[pass+1]; pass++ {
	}
	if pass == 0 || pass == n-1 {
		return false
	}
	for ; pass+1 < n && A[pass] > A[pass+1]; pass++ {
	}
	return pass == n-1
}
