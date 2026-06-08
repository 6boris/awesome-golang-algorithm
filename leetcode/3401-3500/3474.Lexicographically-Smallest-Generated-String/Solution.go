package Solution

func Solution(str1, str2 string) string {
	n, m := len(str1), len(str2)
	s := make([]byte, n+m-1)
	fixed := make([]int, n+m-1)

	for i := range s {
		s[i] = 'a'
	}

	for i := 0; i < n; i++ {
		if str1[i] == 'T' {
			for j := i; j < i+m; j++ {
				if fixed[j] == 1 && s[j] != str2[j-i] {
					return ""
				} else {
					s[j] = str2[j-i]
					fixed[j] = 1
				}
			}
		}
	}

	for i := 0; i < n; i++ {
		if str1[i] == 'F' {
			flag := false
			idx := -1
			for j := i + m - 1; j >= i; j-- {
				if str2[j-i] != s[j] {
					flag = true
				}
				if idx == -1 && fixed[j] == 0 {
					idx = j
				}
			}
			if flag {
				continue
			} else if idx != -1 {
				s[idx] = 'b'
			} else {
				return ""
			}
		}
	}
	return string(s)
}
