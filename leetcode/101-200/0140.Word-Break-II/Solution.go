package Solution

func wordBreak(s string, wordDict []string) []string {
	return dfs(s, wordDict, map[string][]string{})
}

func dfs(s string, wordDict []string, dict map[string][]string) []string {
	if res, ok := dict[s]; ok {
		return res
	}
	if len(s) == 0 {
		return []string{""}
	}

	var res []string

	for _, word := range wordDict {
		if len(word) <= len(s) && word == s[:len(word)] {
			for _, str := range dfs(s[len(word):], wordDict, dict) {
				if len(str) == 0 {
					res = append(res, word)
				} else {
					res = append(res, word+" "+str)
				}
			}
		}
	}
	dict[s] = res
	return res
}
