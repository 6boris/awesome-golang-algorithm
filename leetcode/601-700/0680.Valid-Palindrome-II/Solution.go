package Solution

func Solution(s string) bool {
	size := len(s)
	l, r := 0, size-1
	for ; l < r && s[l] == s[r]; l, r = l+1, r-1 {
	}
	if l >= r {
		return true
	}
	sl, sr := l+1, r
	for ; sl < sr && s[sl] == s[sr]; sl, sr = sl+1, sr-1 {
	}
	if sl >= sr {
		return true
	}

	sl, sr = l, r-1
	for ; sl < sr && s[sl] == s[sr]; sl, sr = sl+1, sr-1 {
	}
	return sl >= sr
}
