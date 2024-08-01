package Solution

import "strconv"

func Solution(details []string) int {
	ans := 0
	for _, d := range details {
		age, _ := strconv.Atoi(d[11:13])
		if age > 60 {
			ans++
		}
	}
	return ans
}
