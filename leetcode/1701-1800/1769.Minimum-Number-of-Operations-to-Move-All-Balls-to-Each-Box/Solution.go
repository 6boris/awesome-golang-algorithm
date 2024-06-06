package Solution

func Solution(boxes string) []int {
	ans := make([]int, len(boxes))
	pos := make([]int, 0)
	for i := 0; i < len(boxes); i++ {
		if boxes[i] == '1' {
			pos = append(pos, i)
		}
	}
	for i := 0; i < len(boxes); i++ {
		for _, p := range pos {
			diff := p - i
			if diff < 0 {
				diff = -diff
			}
			ans[i] += diff
		}
	}
	return ans
}
