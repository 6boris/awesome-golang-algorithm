package Solution

import "fmt"

func countAndSay(n int) string {
	if n < 1 {
		return ""
	}

	say := []byte("1")
	for i := 0; i < n-1; i++ {
		var cnt byte
		var nextSay []byte
		for j := 0; j < len(say); j++ {
			cnt++
			if j == len(say)-1 || say[j] != say[j+1] {
				nextSay = append(nextSay, 48+cnt, say[j]) // 48 - the ASCII code point of "0"
				fmt.Println(string(nextSay), string(48+cnt), string( say[j]))
				cnt = 0
			}
		}
		say = nextSay
	}
	return string(say)
}
