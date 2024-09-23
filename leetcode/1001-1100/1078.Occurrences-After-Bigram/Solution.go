package Solution

import "strings"

func Solution(text string, first string, second string) []string {
	testlist := strings.Split(text, " ")
	fm := make(map[string]struct{})
	sm := make(map[string]struct{})
	for _, w := range strings.Split(first, " ") {
		fm[w] = struct{}{}
	}
	for _, w := range strings.Split(second, " ") {
		sm[w] = struct{}{}
	}
	ans := make([]string, 0)
	for i := 0; i < len(testlist)-2; i++ {
		_, ok1 := fm[testlist[i]]
		_, ok2 := sm[testlist[i+1]]
		if ok1 && ok2 {
			ans = append(ans, testlist[i+2])
		}
	}
	return ans
}
