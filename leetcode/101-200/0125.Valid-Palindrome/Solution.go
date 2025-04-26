package Solution

func Solution(s string) bool {
	var is func(byte) byte
	is = func(b byte) byte {
		if b >= '0' && b <= '9' {
			return b
		}
		if b >= 'a' && b <= 'z' {
			return b
		}
		if b >= 'A' && b <= 'Z' {
			return b + 32
		}
		return byte(' ')
	}
	l, r := 0, len(s)-1
	var bl, br byte
	for l < r {
		for ; l < r; l++ {
			bl = is(s[l])
			if bl != byte(' ') {
				break
			}
		}
		for ; l < r; r-- {
			br = is(s[r])
			if br != byte(' ') {
				break
			}
		}
		if l < r {
			if bl != br {
				return false
			}
			l, r = l+1, r-1
		}
	}
	return true
}
