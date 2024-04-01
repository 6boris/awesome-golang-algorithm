package Solution

func Solution(bookings [][]int, n int) []int {
	ans := make([]int, n)
	for _, booking := range bookings {
		for i := booking[0]; i <= booking[1]; i++ {
			ans[i-1] += booking[2]
		}
	}
	return ans
}
