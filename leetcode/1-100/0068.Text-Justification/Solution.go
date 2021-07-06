package Solution

import "strings"

func Solution(words []string, k int) []string {
	result := make([]string, 0);
	currWords := make([]string, 0);
	currWordsLen := 0;
	for _, word := range words {
		if currWordsLen + len(currWords) + len(word) > k {
			if len(currWords) == 1 {
				result = append(result, currWords[0] + strings.Repeat(" ", k - currWordsLen));
			} else {
				spaces := k - currWordsLen
				spacesBetween, extras := spaces / (len(currWords) - 1), spaces % (len(currWords) - 1);
				for i := 0; i < extras; i++ {
					currWords[i] += " ";
				}
				result = append(result, strings.Join(currWords, strings.Repeat(" ", spacesBetween)));
			}
			currWords = make([]string, 0);
			currWordsLen = 0;
		}
		currWords = append(currWords, word);
		currWordsLen += len(word);
	}
	result = append(result, strings.Join(currWords, " ") + strings.Repeat(" ", k - currWordsLen - len(currWords) + 1));
	return result;
}
