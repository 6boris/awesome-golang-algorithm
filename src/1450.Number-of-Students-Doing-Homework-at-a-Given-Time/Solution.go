package Solution

func Solution(startTime []int, endTime []int, queryTime int) int {
	c := 0
	for i := 0; i < len(startTime); i++ {
		if queryTime >= startTime[i] && queryTime <= endTime[i] {
			c++
		}
	}
	return c
}
