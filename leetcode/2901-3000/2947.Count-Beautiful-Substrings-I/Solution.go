package Solution

func Solution(s string, k int) int {
	var isVowel func(byte) bool
	isVowel = func(b byte) bool {
		return b == 'a' || b == 'e' || b == 'i' || b == 'o' || b == 'u'
	}
	ans := 0
	vc := make([][2]int, len(s)+1)
	for i := 1; i <= len(s); i++ {
		vc[i] = [2]int{vc[i-1][0], vc[i-1][1]}
		if isVowel(s[i-1]) {
			vc[i][0]++
		} else {
			vc[i][1]++
		}
		for j := i - 1; j > 0; j -= 2 {
			a, b := vc[i][0]-vc[j-1][0], vc[i][1]-vc[j-1][1]
			if a == b && (a*b)%k == 0 {
				ans++
			}
		}
	}
	return ans
}
