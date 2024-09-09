package Solution

import "strings"

var skipSymbol = map[byte]struct{}{
	'!': {}, '?': {}, '\'': {}, ',': {}, ';': {}, '.': {},
}

func Solution(paragraph string, banned []string) string {
	words := strings.Split(paragraph, " ")
	bm := make(map[string]struct{})
	for _, b := range banned {
		bm[b] = struct{}{}
	}
	count := make(map[string]int)
	for _, word := range words {
		lower := strings.ToLower(word)
		i := 0
		for ; i < len(word); i++ {
			if _, ok := skipSymbol[word[i]]; ok {
				break
			}
		}
		w := lower[:i]
		if _, ok := bm[w]; ok {
			continue
		}

		count[w]++
	}
	ans := 0
	ansKey := ""
	for k, c := range count {
		if c > ans {
			ansKey = k
			ans = c
		}
	}
	return ansKey
}
