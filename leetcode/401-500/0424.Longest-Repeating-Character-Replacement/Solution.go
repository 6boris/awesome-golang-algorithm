package Solution

func Solution(s string, k int) int {
	byteIndex := make(map[byte]struct{})
	for _, b := range s {
		byteIndex[byte(b)] = struct{}{}
	}
	l := len(s)
	ans := 0
	for b := range byteIndex {
		start, end := 0, 0
		sameByte := 0
		// a b a
		// 0 1 2
		for ; end < l; end++ {
			if s[end] == b {
				sameByte++
			}
			for end+1-start-sameByte > k {
				if s[start] == b {
					sameByte--
				}
				start++
			}
			if end-start+1 > ans {
				ans = end - start + 1
			}
		}
	}
	return ans
}
