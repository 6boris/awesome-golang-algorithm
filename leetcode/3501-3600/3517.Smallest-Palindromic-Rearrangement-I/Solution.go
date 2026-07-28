package Solution

func Solution(s string) string {
	cnt := [26]int{}
	ret := make([]byte, len(s))
	for i := range s {
		cnt[s[i]-'a']++
	}
	oddIndex := -1
	for i := range 26 {
		if cnt[i]&1 == 1 {
			oddIndex = i
			break
		}
	}
	left, right := 0, len(s)-1
	for i := range 26 {
		char := byte(i) + 'a'
		for ; cnt[i] >= 2; cnt[i] -= 2 {
			ret[left], ret[right] = char, char
			left, right = left+1, right-1
		}
	}
	if oddIndex != -1 {
		ret[len(s)/2] = byte(oddIndex) + 'a'
	}
	return string(ret)
}
