package Solution

func Solution(s string) string {
	ret := []byte(s)
	if !(s[0] == '?' && s[1] == '?') {
		if s[0] == '?' {
			ret[0] = '1'
			if s[1] >= '2' {
				ret[0] = '0'
			}
		}
		if s[1] == '?' {
			ret[1] = '9'
			if s[0] == '1' {
				ret[1] = '1'
			}
		}
	} else {
		ret[0], ret[1] = '1', '1'
	}
	if s[3] == '?' {
		ret[3] = '5'
	}
	if s[4] == '?' {
		ret[4] = '9'
	}
	return string(ret)
}
