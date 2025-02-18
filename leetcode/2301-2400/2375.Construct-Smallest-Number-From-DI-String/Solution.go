package Solution

func Solution(pattern string) string {
	res := make([]byte, len(pattern)+1)
	for i := range res {
		res[i] = '0'
	}

	index := byte('1')
	for i := 0; i < len(pattern); i++ {
		if pattern[i] == 'D' {
			continue
		}
		res[i] = index
		index++
		for pre := i - 1; pre >= 0 && res[pre] == '0'; pre-- {
			res[pre] = index
			index++
		}
	}

	res[len(pattern)] = index
	index++
	for pre := len(pattern) - 1; pre >= 0 && res[pre] == '0'; pre-- {
		res[pre] = index
		index++
	}
	return string(res)
}
