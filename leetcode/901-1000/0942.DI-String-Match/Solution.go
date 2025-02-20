package Solution

func Solution(s string) []int {
	res := make([]int, len(s)+1)
	for i := range res {
		res[i] = -1
	}

	index := 0
	for i := 0; i < len(s); i++ {
		if s[i] == 'D' {
			continue
		}
		res[i] = index
		index++
		for pre := i - 1; pre >= 0 && res[pre] == -1; pre-- {
			res[pre] = index
			index++
		}
	}
	res[len(s)] = index
	index++
	for pre := len(s) - 1; pre >= 0 && res[pre] == -1; pre-- {
		res[pre] = index
		index++
	}
	return res
}
