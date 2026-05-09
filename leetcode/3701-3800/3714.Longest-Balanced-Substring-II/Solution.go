package Solution

func longestPairBalanced(s string, x, y, z byte) int {
	ans, xCount, yCount := 0, 0, 0
	mp := map[int]int{0: -1} // diff -> index

	for i := 0; i < len(s); i++ {
		if s[i] == z {
			mp = map[int]int{0: i}
			xCount, yCount = 0, 0
			continue
		}
		if s[i] == x {
			xCount++
		} else if s[i] == y {
			yCount++
		}
		diff := xCount - yCount
		if firstIdx, ok := mp[diff]; ok {
			ans = max(ans, i-firstIdx)
		} else {
			mp[diff] = i
		}
	}
	return ans
}

func Solution(s string) int {
	n := len(s)
	if n == 0 {
		return 0
	}
	ans := 0

	// CASE 1: 1种字符 (最长连续段)
	tempFre := 0
	for i := 0; i < n; i++ {
		if i > 0 && s[i] == s[i-1] {
			tempFre++
		} else {
			tempFre = 1
		}
		if tempFre > ans {
			ans = tempFre
		}
	}

	// CASE 2: 2种字符
	ans = max(ans, longestPairBalanced(s, 'a', 'b', 'c'))
	ans = max(ans, longestPairBalanced(s, 'a', 'c', 'b'))
	ans = max(ans, longestPairBalanced(s, 'b', 'c', 'a'))

	// CASE 3: 3种字符 (State Hash: a-b, b-c)
	type State struct{ d1, d2 int }
	mp := map[State]int{{0, 0}: -1}
	a, b, c := 0, 0, 0
	last := [3]int{-1, -1, -1} // a, b, c 最近出现位置

	for i := 0; i < n; i++ {
		if s[i] == 'a' {
			a++
			last[0] = i
		}
		if s[i] == 'b' {
			b++
			last[1] = i
		}
		if s[i] == 'c' {
			c++
			last[2] = i
		}

		st := State{a - b, b - c}
		if firstIdx, ok := mp[st]; ok {
			// 必须保证 a,b,c 在此区间内都出现过
			if last[0] > firstIdx && last[1] > firstIdx && last[2] > firstIdx {
				if i-firstIdx > ans {
					ans = i - firstIdx
				}
			}
		} else {
			mp[st] = i
		}
	}
	return ans
}
