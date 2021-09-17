package Solution

import (
	"strings"
)

func canConstruct(ransomNote string, magazine string) bool {
	for _, c := range ransomNote {
		idx := strings.IndexRune(magazine, c)
		if idx > -1 {
			magazine = magazine[:idx] + magazine[idx+1:]
		} else {
			return false
		}
	}
	return true
}

func canConstruct2(ransomNote string, magazine string) bool {
	m := make(map[rune]int)

	for _, v := range magazine {
		m[v]++
	}
	for _, v := range ransomNote {
		if m[v] == 0 {
			return false
		}
		m[v]--
	}
	return true
}
