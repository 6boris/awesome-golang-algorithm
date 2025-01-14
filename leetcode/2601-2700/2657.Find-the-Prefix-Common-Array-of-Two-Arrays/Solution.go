package Solution

func Solution(A []int, B []int) []int {
	ans := make([]int, len(A))
	count := make([]uint8, len(A)+1)
	for i := 0; i < len(A); i++ {
		if i > 0 {
			ans[i] = ans[i-1]
		}
		count[A[i]]++
		count[B[i]]++
		if A[i] == B[i] {
			ans[i]++
			continue
		}
		if count[A[i]] == 2 {
			ans[i]++
		}
		if count[B[i]] == 2 {
			ans[i]++
		}
	}
	return ans
}
