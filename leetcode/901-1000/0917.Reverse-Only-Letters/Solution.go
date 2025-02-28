package Solution

func ok(b byte) bool {
	return (b >= 'a' && b <= 'z') || (b >= 'A' && b <= 'Z')
}

func Solution(s string) string {
	bs := []byte(s)
	l, r := 0, len(bs)-1
	for l < r {
		for ; l < r && !ok(bs[l]); l++ {

		}
		for ; r > l && !ok(bs[r]); r-- {
		}
		bs[l], bs[r] = bs[r], bs[l]
		l, r = l+1, r-1
	}
	return string(bs)
}
