package Solution

import "strings"

// Eating, Eating right now
func match1813(s, t []string) bool {
	l1, r1 := 0, len(s)-1
	l2, r2 := 0, len(t)-1
	for ; l1 <= r1 && s[l1] == t[l2]; l1, l2 = l1+1, l2+1 {

	}
	// 前缀
	if l1 > r1 {
		return true
	}

	for ; r1 >= l1 && s[r1] == t[r2]; r1, r2 = r1-1, r2-1 {
	}
	if r1 < l1 {
		return true
	}
	return false
}

func Solution(sentence1 string, sentence2 string) bool {
	// 相等只需要插入空的数据即可
	if sentence1 == sentence2 {
		return true
	}
	l1, l2 := len(sentence1), len(sentence2)
	// 如果长度相等，但是字符串并不相等，说明无法插入
	if l1 == l2 {
		return false
	}
	s1, s2 := strings.Split(sentence1, " "), strings.Split(sentence2, " ")
	if l1 < l2 {
		return match1813(s1, s2)
	}
	return match1813(s2, s1)

}
