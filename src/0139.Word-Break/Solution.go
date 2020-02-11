package Solution

//	暴力遍历 比较 时间复杂度到 O(N3)
func wordBreak(s string, wordDict []string) bool {
	dp := make([]bool, len(s)+1)
	dp[0] = true

	for i := 1; i <= len(s); i++ {
		for j := 0; j <= i-1; j++ {
			if dp[j] == true {
				for _, word := range wordDict {
					if s[j:i] == word {
						dp[i] = true
						break
					}
				}
			}

		}
	}
	return dp[len(s)]
}

//	一次遍历时间复杂度 O(N2)
func wordBreak2(s string, wordDict []string) bool {
	dp := make([]bool, len(s)+1)
	dp[0] = true

	for i := 0; i < len(dp); i++ {
		if dp[i] {
			for _, word := range wordDict {
				//	巧妙的运用 i + word len 避免2层循环
				if i+len(word) <= len(s) && s[i:i+len(word)] == word {
					dp[i+len(word)] = true
				}
			}
		}
	}
	return dp[len(s)]
}

//	递归遍历
func wordBreak3(s string, wordDict []string) bool {
	dict := make(map[string]bool)
	return dfs(s, wordDict, dict)
}

//	s 每次比对的字符串
//	wordDict 目标字典列表
// 	匹配结果
func dfs(s string, wordDict []string, dict map[string]bool) bool {
	//	终止条件
	if len(s) == 0 {
		return true
	}

	if res, ok := dict[s]; ok {
		return res
	}

	for _, word := range wordDict {
		if len(word) > len(s) || word != s[:len(word)] {
			continue
		}

		if dfs(s[len(word):], wordDict, dict) {
			dict[s] = true
			return true
		}

	}
	dict[s] = false

	return false
}
