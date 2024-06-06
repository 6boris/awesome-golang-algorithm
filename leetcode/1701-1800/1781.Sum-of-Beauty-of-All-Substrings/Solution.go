package Solution

func Solution(s string) int {
	ans := 0
	//ab,aa
	for end := 2; end < len(s); end++ {
		tmp := [26]int{}
		tmp[s[end]-'a']++
		a := 0
		for pre := end - 1; pre >= 0; pre-- {
			b := 0
			tmp[s[pre]-'a']++
			a = max(a, tmp[s[pre]-'a'])
			for i := 0; i < 26; i++ {
				if tmp[i] != 0 {
					if b == 0 || b > tmp[i] {
						b = tmp[i]
					}
				}
			}
			ans += a - b
		}
	}
	return ans
}
