package Solution

func Solution(A []int) bool {
	inc, dec := true, true
	for i := 1; i < len(A); i++ {
		if A[i] > A[i-1] {
			dec = false
		} else if A[i] < A[i-1] {
			inc = false
		}
	}
	return inc || dec
}
