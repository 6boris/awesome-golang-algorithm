package Solution

func Solution(lcp [][]int) string {
	n := len(lcp)
	word := make([]byte, n)
	current := byte('a')

	for i := 0; i < n; i++ {
		if word[i] == 0 {
			if current > 'z' {
				return ""
			}
			word[i] = current
			for j := i + 1; j < n; j++ {
				if lcp[i][j] > 0 {
					word[j] = word[i]
				}
			}
			current++
		}
	}

	// verify if the constructed string meets the LCP matrix requirements
	for i := n - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			if word[i] != word[j] {
				if lcp[i][j] != 0 {
					return ""
				}
			} else {
				if i == n-1 || j == n-1 {
					if lcp[i][j] != 1 {
						return ""
					}
				} else {
					if lcp[i][j] != lcp[i+1][j+1]+1 {
						return ""
					}
				}
			}
		}
	}

	return string(word)
}
