package Solution

const mod1639 = 1000000007

func Solution(words []string, target string) int {
	lt := len(target)
	lw := len(words[0])
	dp := make([][]int, lt+1)
	for i := 0; i <= lt; i++ {
		dp[i] = make([]int, lw+1)
	}
	wordsMap := make([]map[byte][]int, len(words))
	for i, word := range words {
		wordsMap[i] = make(map[byte][]int)
		for idx, b := range word {
			if _, ok := wordsMap[i][byte(b)]; !ok {
				wordsMap[i][byte(b)] = make([]int, 0)
			}
			wordsMap[i][byte(b)] = append(wordsMap[i][byte(b)], idx+1)
		}
	}

	for index := 1; index <= lt; index++ {
		addOne := 0
		if index == 1 {
			addOne = 1
		}
		for _, word := range wordsMap {
			if v, ok := word[target[index-1]]; ok {
				for _, i := range v {
					dp[index][i] = (dp[index][i] + dp[index-1][i-1] + addOne) % mod1639
				}
			}
		}
		for i := 2; i <= lw; i++ {
			dp[index][i] = (dp[index][i] + dp[index][i-1]) % mod1639
		}
	}
	return dp[lt][lw]
}
