package Solution

func Solution(n int) int {
	sum := make([]int, n+1)

	for i := 1; i <= n; i++ {
		sum[i] = sum[i-1] + i
	}
	// 1, 2, 3, 4, 5,  6, 7, 8
	// 1, 3, 6,10, 15,21, 28,36
	left, right := 1, n
	for left <= right {
		mid := (right-left)/2 + left
		a := sum[mid]
		b := sum[n] - sum[mid-1]
		if a == b {
			return mid
		}
		if a > b {
			right = mid - 1
			continue
		}
		left = mid + 1
	}
	return -1
}
