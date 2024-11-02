package Solution

func Solution(sentence string) bool {
	l := len(sentence)
	if sentence[0] != sentence[l-1] {
		return false
	}
	for i := range sentence {
		if sentence[i] == ' ' {
			if sentence[i-1] != sentence[i+1] {
				return false
			}
		}
	}
	return true
}
