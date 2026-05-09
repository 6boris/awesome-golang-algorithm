package Solution

import "strings"

func Solution(sentence string) int {
	words := strings.Split(sentence, " ")
	ret := 0
	validate := func(str string) bool {
		size := len(str)
		if size == 0 {
			return false
		}
		foundHyphens := false
		for i := range size {
			if str[i] >= '0' && str[i] <= '9' {
				// 包含数字
				return false
			}
			if (str[i] == '.' || str[i] == '!' || str[i] == ',') && i != size-1 {
				return false
			}
			if str[i] == '-' {
				if foundHyphens {
					return false
				}
				foundHyphens = true
				if !(i > 0 && i < size-1 && str[i-1] >= 'a' && str[i-1] <= 'z' && str[i+1] >= 'a' && str[i+1] <= 'z') {
					return false
				}
			}
		}
		return true
	}
	for i := range words {
		if validate(words[i]) {
			ret++
		}
	}
	return ret
}
