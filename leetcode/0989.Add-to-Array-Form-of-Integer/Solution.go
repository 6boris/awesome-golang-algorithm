package Solution

func Solution(A []int, K int) []int {
	var res []int
	for i := len(A) - 1; i >= 0; i-- {
		res = append(res, (A[i]+K)%10)
		K = (A[i] + K) / 10
	}
	for K > 0 {
		res = append(res, K%10)
		K /= 10
	}
	return reverse(res)
}

func reverse(arr []int) []int {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
	return arr
}
