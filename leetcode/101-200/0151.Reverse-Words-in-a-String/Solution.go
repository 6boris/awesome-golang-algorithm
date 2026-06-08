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
