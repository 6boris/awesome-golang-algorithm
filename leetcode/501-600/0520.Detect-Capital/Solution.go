package Solution

func Solution(word string) bool {
	length := len(word)
	if length <= 1 {
		return true
	}
	cap := false
	if word[0] >= 65 && word[0] <= 90 {
		cap = true
	}

	haveCap, haveLow := false, false
	for idx := 1; idx < length; idx++ {
		if word[idx] >= 65 && word[idx] <= 90 {
			haveCap = true
		} else {
			haveLow = true
		}
		if haveCap && haveLow {
			break
		}
	}
	if cap && !(haveCap && haveLow) {
		return true
	}
	return haveLow && !haveCap
}
