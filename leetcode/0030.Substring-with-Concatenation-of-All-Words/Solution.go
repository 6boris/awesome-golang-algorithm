package Solution

func findSubstring(s string, words []string) []int {
	//init
	res := []int{}

	sLen, wordsNum := len(s), len(words)
	//special
	if sLen == 0 || wordsNum == 0 {
		return res
	}
	//注意：words中所有单词长度相同，均为wordLen
	wordLen := len(words[0])
	// remainNum 记录 words 中每个单词还能出现的次数
	remainNum := make(map[string]int, wordsNum)
	count := 0

	// 由于 index 增加的跨度是 wordLen
	// index 需要分别从{0,1,...,wordLen-1}这些位置开始检查，才能不遗漏
	for initialIndex := 0; initialIndex < wordLen; initialIndex++ {
		index := initialIndex
		count = initRecord(words, remainNum)

		for index+wordsNum*wordLen <= sLen {
			word := s[index+count*wordLen : index+(count+1)*wordLen]
			remainTimes, ok := remainNum[word]
			switch {
			case !ok:
				//出现了不在words中的单词
				//从word后面一个单词，重新开始统计
				index += wordLen * (count + 1)
				if count != 0 {
					count = initRecord(words, remainNum)
				}
			case remainTimes == 0:
				//word的出现次数上次查询就用完了
				//说明s[index:index+(count+1)*wordLen]中有单词多出现了
				index, count = moveIndex(index, wordLen, count, s, remainNum)
				// 这个case会连续出现
				// 直到s[index:index+(count+1)*lenw]中所有单词的出现次数都不超标
				//
				// 在 moveIndex() 的过程中，index+(count+1)*lenw 保持值不变
			default:
				// ok && remainTimes > 0，word 符合出现的条件
				// 更新统计记录
				remainNum[word]--
				count++
				// 检查 words 能否排列组合成 s[index:index+count*lenw]
				if count == wordsNum {
					res = append(res, index)

					// 把 index 指向下一个单词
					// 开始下一个统计
					index, count = moveIndex(index, wordLen, count, s, remainNum)
				}
			}
		}
	}
	return res
}

func initRecord(words []string, remainNum map[string]int) int {
	for _, word := range words {
		remainNum[word] = 0
	}
	for _, word := range words {
		remainNum[word]++
	}
	return 0
}

func moveIndex(index, wordLen, count int, s string, remainNum map[string]int) (int, int) {
	// index 指向下一个单词的同时，需要同时修改统计记录

	// 增加 index 指向的 word 可出现次数一次，
	remainNum[s[index:index+wordLen]]++
	// 连续符合条件的单词数减少一个
	count--
	// index 后移一个 word 的长度
	index += wordLen
	return index, count
}
