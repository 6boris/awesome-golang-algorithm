package Solution

func Solution(word string) int {
	a, b := [26]int{}, [26]int{}
	for i := range word {
		if word[i] >= 'a' && word[i] <= 'z' {
			a[word[i]-'a']++
			continue
		}
		b[word[i]-'A']++
	}
	var ret int
	for i := range 26 {
		if a[i] > 0 && b[i] > 0 {
			ret++
		}
	}
	return ret
}
