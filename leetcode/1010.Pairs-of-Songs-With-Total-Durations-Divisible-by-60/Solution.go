package Solution

func Solution(time []int) int {
	ans, hash := 0, make([]int, 60)
	for _, t := range time {
		v := t % 60
		if v == 0 {
			ans += hash[0]
		} else {
			ans += hash[60-v]
		}
		hash[v]++
	}
	return ans
}
