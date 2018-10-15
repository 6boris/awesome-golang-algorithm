package Solution

// O(n) time O(1) space Solution
func lengthOfLongestSubstring(s string) int {
	var chPosition [256]int // [0, 0, 0, ...]
	maxLength, substringLen, lastRepeatPos := 0, 0, 0

	for i := 0; i < len(s); i++ {
		if pos := chPosition[s[i]]; pos > 0 {
			// record current substring length
			maxLength = Max(substringLen, maxLength)

			// update characters position
			chPosition[s[i]] = i + 1

			// update last repeat character position
			lastRepeatPos = Max(pos, lastRepeatPos)

			// update the current substring from last repeat character
			substringLen = i + 1 - lastRepeatPos
		} else {
			substringLen += 1
			chPosition[s[i]] = i + 1
		}
	}

	return Max(maxLength, substringLen)
}

//	暴力求解(会超时)
func lengthOfLongestSubstring2(s string) int {
	ans := 0

	for i := 0; i < len(s); i++ {
		for j := i + 1; j <= len(s); j++ {
			if allUnique(s, i, j) {
				ans = Max(ans, j-i)
			}
		}
	}
	return ans
}

func allUnique(s string, start int, end int) bool {
	sMap := make(map[string]int)

	for i := start; i < end; i++ {
		if sMap[string(s[i])] > 0 {
			return false
		}
		sMap[string(s[i])]++
	}

	return true
}

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
