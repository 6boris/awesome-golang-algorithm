package Solution

import (
	"strings"
)

func lengthOfLastWord(s string) int {
	start := len(s) - 1
	//	寻找最后一个单词的尾下标
	for start >= 0 && s[start] == ' ' {
		start--
	}
	//	寻找最后一个单词的初始下标
	end := start
	for start >= 0 && s[start] != ' ' {
		start--
	}

	return end - start
}

func lengthOfLastWord1(s string) int {
	s = strings.TrimRight(s, " ")
	return len(s) - strings.LastIndex(s, " ") - 1
}
