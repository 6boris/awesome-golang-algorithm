package Solution

import "fmt"

func minWindow(s string, t string) string {
	hash := make(map[uint8]int)
	cut := 0

	for i := 0; i < len(s); i++ {
		fmt.Println(s[i])
		if _, ok := hash[s[i]]; !ok {
			cut++
		}
		hash[s[i]]++
	}

	res := ""
	for i, j, c := 0, 0, 0; i < len(s); i++ {
		if hash[s[i]] == 1 {
			c++
		}
		hash[s[i]]--
		for c == cut && hash[s[i]] < 0 {
			hash[s[j]]++
			j++
		}
		if c == cut {
			if res == "" || len(res) > i-j+1 {
				res = s[j : i-j+1]
			}
		}

	}
	return res
}
