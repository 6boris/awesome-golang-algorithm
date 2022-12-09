package Solution

import "sort"

func Solution(s string) string {

	sb := []byte(s)
	byteCount := make(map[byte]int)
	for _, b := range sb {
		byteCount[b]++
	}
	sort.Slice(sb, func(i, j int) bool {
		if byteCount[sb[i]] == byteCount[sb[j]] {
			return sb[i] < sb[j]
		}
		return byteCount[sb[i]] > byteCount[sb[j]]
	})
	// z知道每个字符的出现次数, 然后排序，写入
	return string(sb)
}
