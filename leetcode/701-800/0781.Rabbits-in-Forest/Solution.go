package Solution

func Solution(answers []int) int {
	count := make(map[int]int)
	for _, say := range answers {
		count[say]++
	}
	ans := 0
	for say, c := range count {
		if say == 0 {
			ans += c
			continue
		}

		need := say + 1
		diff := c / need
		if c%need != 0 {
			diff++
		}
		ans += diff * need
	}
	return ans
}
