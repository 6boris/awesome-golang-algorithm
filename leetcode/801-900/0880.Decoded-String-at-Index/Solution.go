package Solution

import "fmt"

func Solution(s string, k int) string {
	size := uint64(0)
	n := len(s)
	for i := 0; i < n; i++ {
		if s[i] >= 'a' && s[i] <= 'z' {
			size++
			continue
		}
		size = size * uint64(s[i]-'0')
	}
	kk := uint64(k)
	for i := n - 1; i >= 0; i-- {
		c := s[i]
		kk %= size
		if kk == 0 && c >= 'a' && c <= 'z' {
			return fmt.Sprintf("%c", c)
		}
		if c >= 'a' && c <= 'z' {
			size--
			continue
		}
		size = size / uint64(c-'0')
	}
	return ""
}
