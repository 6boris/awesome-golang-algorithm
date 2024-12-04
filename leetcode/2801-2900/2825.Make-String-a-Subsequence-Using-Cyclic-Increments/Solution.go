package Solution

func Solution(str1 string, str2 string) bool {
	if len(str2) > len(str1) {
		return false
	}
	b := 0
	for i := 0; i < len(str1) && b < len(str2); i++ {
		t1 := str1[i] - 'a'
		t2 := str2[b] - 'a'
		if t1 == t2 || t2 == (t1+1+26)%26 {
			b++
		}
	}
	return b == len(str2)
}
