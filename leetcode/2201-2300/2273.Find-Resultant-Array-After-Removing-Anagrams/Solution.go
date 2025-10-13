package Solution

func equal(a, b [26]int) bool {
	for i := range 26 {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func Solution(words []string) []string {
	l := len(words)
	indies := make([][26]int, l)
	for index, word := range words {
		for _, b := range word {
			indies[index][b-'a']++
		}
	}
	index := 0
	stack := make([]int, l)
	for i := 1; i < l; i++ {
		if equal(indies[stack[index]], indies[i]) {
			continue
		}
		index++
		stack[index] = i
	}
	var ret []string
	for i := range index + 1 {
		ret = append(ret, words[stack[i]])
	}
	return ret
}
