package Solution

func Solution(answerKey string, k int) int {
	count := make(map[byte]int)
	for i := 0; i < k; i++ {
		count[answerKey[i]]++
	}
	ans := k
	left := 0
	for right := k; right < len(answerKey); right++ {
		count[answerKey[right]]++
		for {
			a := count['T']
			if count['F'] < a {
				a = count['F']
			}
			if a <= k {
				break
			}
			count[answerKey[left]]--
			left++
		}
		if r := right - left + 1; r > ans {
			ans = r
		}
	}
	return ans
}
