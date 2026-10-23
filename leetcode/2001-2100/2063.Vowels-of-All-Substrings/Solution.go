package Solution

func Solution(word string) int64 {
	n := int64(len(word))
	var ret int64
	for i := int64(0); i < n; i++ {
		if word[i] == 'a' || word[i] == 'e' || word[i] == 'i' || word[i] == 'o' || word[i] == 'u' {
			prev := i
			next := n - int64(i) - 1
			ret += prev + next + prev*next + 1
		}
	}
	return ret
}
