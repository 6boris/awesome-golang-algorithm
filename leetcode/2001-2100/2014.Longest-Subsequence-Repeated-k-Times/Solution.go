package Solution

func Solution(s string, k int) string {
	count := [26]int{}
	for _, b := range s {
		count[b-'a']++
	}
	var greaterK []string
	for i := 25; i >= 0; i-- {
		if count[i] >= k {
			greaterK = append(greaterK, string('a'+byte(i)))
		}
	}

	var match func(string) bool
	match = func(pattern string) bool {
		i, c := 0, 0
		for j := 0; j < len(s); j++ {
			if s[j] == pattern[i] {
				i++
				if i == len(pattern) {
					c++
					i = 0
					if c >= k {
						return true
					}
				}
			}
		}
		return false
	}

	loop := make([]string, len(greaterK))
	copy(loop, greaterK)
	// z, y, x, 我们应该需要判断 z ok (zz) ok , zzz(ok) / 以及 zx, zy ...
	var ans string
	for len(loop) > 0 {
		cur := loop[0]
		loop = loop[1:]
		if len(cur) > len(ans) {
			ans = cur
		}
		// 开始做组合
		for _, p := range greaterK {
			pattern := cur + p
			if match(pattern) {
				// 这个匹配后，我们需要判断是否还可以继续与其他的z,y x这种组合
				loop = append(loop, pattern)
			}
		}
	}

	return ans
}
