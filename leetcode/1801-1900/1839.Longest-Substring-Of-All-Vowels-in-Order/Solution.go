package Solution

type charCount struct {
	char  byte
	count int
}

func Solution(word string) int {

	pre := byte(' ')
	// a, e, i, o, u
	group := make([]charCount, 0)
	index := -1
	for _, b := range []byte(word) {
		if b != pre {
			group = append(group, charCount{b, 1})
			index++
			pre = b
			continue
		}
		group[index].count++
	}
	ret := 0
	// 1,2,3,4,5
	for i := 0; i <= len(group)-5; i++ {
		if group[i].char == 'a' && group[i+1].char == 'e' && group[i+2].char == 'i' && group[i+3].char == 'o' && group[i+4].char == 'u' {
			ret = max(ret, group[i].count+group[i+1].count+group[i+2].count+group[i+3].count+group[i+4].count)
		}
	}
	return ret
}
