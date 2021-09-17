package Solution

import (
	"strings"
)

func reverseWords(s string) string {
	//	用自带的函数拆分数组
	words := strings.Fields(s)

	i := 0
	j := len(words) - 1

	for i < j {
		//	左右22交换
		words[i], words[j] = words[j], words[i]
		i++
		j--
	}

	return strings.Join(words, " ")
}

func reverseWords2(s string) string {
	words := strings.Fields(s)
	return strings.Join(reverseSlice(words), " ")
}

// Helper func that reverses the elements of a string slice.
func reverseSlice(s []string) []string {
	size := len(s)
	for i := range s[:size/2] {
		s[i], s[size-1-i] = s[size-1-i], s[i]
	}
	return s
}
