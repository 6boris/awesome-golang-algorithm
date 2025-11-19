package Solution

func Solution(s string, a int, b int) string {
	n := len(s)
	vis := make([]bool, n)
	res := s
	s = s + s
	for i := 0; !vis[i]; i = (i + b) % n {
		vis[i] = true
		for j := 0; j < 10; j++ {
			kLimit := 0
			if b%2 != 0 {
				kLimit = 9
			}
			for k := 0; k <= kLimit; k++ {
				// before each accumulation, re-truncate t
				t := []byte(s[i : i+n])
				for p := 1; p < n; p += 2 {
					t[p] = '0' + byte((int(t[p]-'0')+j*a)%10)
				}
				for p := 0; p < n; p += 2 {
					t[p] = '0' + byte((int(t[p]-'0')+k*a)%10)
				}
				tStr := string(t)
				if tStr < res {
					res = tStr
				}
			}
		}
	}
	return res
}
