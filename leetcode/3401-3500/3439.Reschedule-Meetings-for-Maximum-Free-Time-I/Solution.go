package Solution

func Solution(eventTime int, k int, startTime []int, endTime []int) int {
	n := len(startTime)
	res := 0
	sum := make([]int, n+1)
	for i := 0; i < n; i++ {
		sum[i+1] = sum[i] + endTime[i] - startTime[i]
	}
	for i := k - 1; i < n; i++ {
		var right int
		if i == n-1 {
			right = eventTime
		} else {
			right = startTime[i+1]
		}
		var left int
		if i == k-1 {
			left = 0
		} else {
			left = endTime[i-k]
		}
		res = max(res, right-left-(sum[i+1]-sum[i-k+1]))
	}
	return res
}
