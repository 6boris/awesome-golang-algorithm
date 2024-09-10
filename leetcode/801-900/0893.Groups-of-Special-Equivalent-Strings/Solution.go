package Solution

func Solution(words []string) int {
	in := make(map[[2][26]int]int)
	for _, word := range words {
		k1, k2 := [26]int{}, [26]int{}
		for i := range word {
			if i&1 == 0 {
				k1[word[i]-'a']++
				continue
			}
			k2[word[i]-'a']++
		}
		in[[2][26]int{k1, k2}]++
	}
	return len(in)
}
