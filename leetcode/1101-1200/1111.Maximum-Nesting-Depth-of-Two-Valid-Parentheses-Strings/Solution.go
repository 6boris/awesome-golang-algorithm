package Solution

func Solution(seq string) []int {
	l := len(seq)
	ret := make([]int, l)

	depth := 0
	bs := []byte(seq)
	for i := range bs {
		if bs[i] == '(' {
			depth++
			ret[i] = 1 - depth&1
			continue
		}
		ret[i] = 1 - depth&1
		depth--
	}
	return ret
}
