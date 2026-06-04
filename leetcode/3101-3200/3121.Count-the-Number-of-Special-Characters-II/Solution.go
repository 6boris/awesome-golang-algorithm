package Solution

func Solution(word string) int {
	lower, upper := [26]int{}, [26]int{}
	for i := range 26 {
		lower[i] = -1
		upper[i] = -1
	}

	for i := range word {
		if word[i] >= 'a' && word[i] <= 'z' {
			lower[word[i]-'a'] = i
			continue
		}
		if upper[word[i]-'A'] == -1 {
			upper[word[i]-'A'] = i
		}
	}

	var ret int
	for i := range 26 {
		if lower[i] != -1 && upper[i] != -1 && lower[i] < upper[i] {
			ret++
		}
	}
	return ret
}
