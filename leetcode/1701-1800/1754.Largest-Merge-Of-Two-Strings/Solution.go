package Solution

import "bytes"

func Solution(word1 string, word2 string) string {
	buf := bytes.Buffer{}
	i, j := 0, 0

	for i < len(word1) && j < len(word2) {
		selectI := false
		if word1[i] == word2[j] {
			// 相等情况
			ii, jj := i+1, j+1
			for ; ii < len(word1) && jj < len(word2) && word1[ii] == word2[jj]; ii, jj = ii+1, jj+1 {
			}
			if ii == len(word1) {
				selectI = false
			} else if jj == len(word2) {
				selectI = true
			} else {
				selectI = word1[ii] > word2[jj]
			}
		}

		if selectI || word1[i] > word2[j] {
			buf.WriteByte(word1[i])
			i++
			continue
		}
		if !selectI || word1[i] < word2[j] {
			buf.WriteByte(word2[j])
			j++
			continue
		}
	}
	for ; i < len(word1); i++ {
		buf.WriteByte(word1[i])
	}
	for ; j < len(word2); j++ {
		buf.WriteByte(word2[j])
	}
	return buf.String()
}
