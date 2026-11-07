package Solution

import (
	"bytes"
	"strings"
)

func Solution(caption string) string {
	src := strings.TrimSpace(caption)
	b := []byte(src)
	got := bytes.Split(b, []byte{' '})
	ret := make([]byte, 100)
	ret[0] = '#'
	var (
		tmp   []byte
		index = 1
	)
	for i := range got {
		tmp = bytes.ToLower(got[i])
		if len(tmp) == 0 {
			continue
		}
		if i != 0 {
			tmp[0] -= 32
		}
		for j := 0; j < len(tmp) && index < 100; j, index = j+1, index+1 {
			if (tmp[j] >= 'a' && tmp[j] <= 'z') || (tmp[j] >= 'A' && tmp[j] <= 'Z') {
				ret[index] = tmp[j]
			}
		}
	}
	return string(ret[:index])
}
