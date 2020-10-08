package Solution

func Solution(A []int, queries [][]int) []int {
	res, evenSum := make([]int, 0), 0
	for _, val := range A {
		if val%2 == 0 {
			evenSum += val
		}
	}
	for _, query := range queries {
		val, index := query[0], query[1]
		if A[index]%2 == 0 {
			evenSum -= A[index]
		}
		A[index] += val
		if A[index]%2 == 0 {
			evenSum += A[index]
		}
		res = append(res, evenSum)
	}
	return res
}
