package Solution

func Solution(s string) int {
	ans := -1
	cache := [26]map[int]int{}
	for i := range 26 {
		cache[i] = make(map[int]int)
	}
	l := len(s)
	start := 0
	cur := s[start]
	for i := 1; i < l; i++ {
		if s[i] == cur {
			continue
		}

		index := cur - 'a'
		count := i - start
		for ll := 1; ll <= count; ll++ {
			cache[index][ll] += count - ll + 1
			if cache[index][ll] >= 3 && ll > ans {
				ans = ll
			}
		}
		cur = s[i]
		start = i
	}
	index := cur - 'a'
	count := l - start
	for ll := 1; ll <= count; ll++ {
		cache[index][ll] += count - ll + 1
		if cache[index][ll] >= 3 && ll > ans {
			ans = ll
		}
	}

	return ans
}
