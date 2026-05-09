package Solution

import "bytes"

func Solution(s string) string {
	parts := bytes.Split([]byte(s), []byte{' '})

	cmp := 0
	for idx := range parts[0] {
		if parts[0][idx] == 'a' || parts[0][idx] == 'e' || parts[0][idx] == 'i' || parts[0][idx] == 'o' || parts[0][idx] == 'u' {
			cmp++
		}
	}
	// if cmp == 0 {
	// 	return s
	// }
	buf := make([][]byte, len(parts))
	buf[0] = parts[0]

	cnt := 0
	for i := 1; i < len(parts); i++ {
		cnt = 0
		for idx := range parts[i] {
			if parts[i][idx] == 'a' || parts[i][idx] == 'e' || parts[i][idx] == 'i' || parts[i][idx] == 'o' || parts[i][idx] == 'u' {
				cnt++
			}
		}
		if cnt == cmp {
			for s, e := 0, len(parts[i])-1; s < e; s, e = s+1, e-1 {
				parts[i][s], parts[i][e] = parts[i][e], parts[i][s]
			}
		}
		buf[i] = parts[i]
	}
	return string(bytes.Join(buf, []byte{' '}))
}
