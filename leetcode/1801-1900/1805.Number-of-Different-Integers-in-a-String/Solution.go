package Solution

func Solution(word string) int {
	count := make(map[string]struct{})
	firstMeet := true
	foundNumber := false
	start := -1
	for i, b := range word {
		if b >= '0' && b <= '9' {
			foundNumber = true
			if firstMeet && b == '0' {
				continue
			}
			if start == -1 {
				start = i
			}
			firstMeet = false
			continue
		}
		if foundNumber {
			if start == -1 {
				count["0"] = struct{}{}
			} else {
				count[word[start:i]] = struct{}{}
			}
		}
		foundNumber = false
		firstMeet = true
		start = -1
	}
	if foundNumber {
		if start == -1 {
			count["0"] = struct{}{}
		} else {
			count[word[start:]] = struct{}{}
		}
	}
	return len(count)
}
