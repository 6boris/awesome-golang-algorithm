package Solution

import "strings"

func Solution(words []string) int {
	ans := 0
	for i := 0; i < len(words); i++ {
		for j := i + 1; j < len(words); j++ {
			if strings.HasPrefix(words[j], words[i]) && strings.HasSuffix(words[j], words[i]) {
				ans++
			}
		}
	}
	return ans
}
