package Solution

import "sort"

func isVowel2785(b byte) bool {
	return b == 'a' || b == 'e' || b == 'i' || b == 'o' || b == 'u' ||
		b == 'A' || b == 'E' || b == 'I' || b == 'O' || b == 'U'
}

func Solution(s string) string {
	vowels := make([]byte, 0)
	indies := make([]int, 0)
	t := make([]byte, len(s))
	for idx, b := range []byte(s) {
		if isVowel2785(b) {
			vowels = append(vowels, b)
			indies = append(indies, idx)
			continue
		}
		t[idx] = b
	}
	sort.Slice(vowels, func(i, j int) bool {
		return vowels[i] < vowels[j]
	})
	for idx, targetIdx := range indies {
		t[targetIdx] = vowels[idx]
	}

	return string(t)
}
