package Solution

func Solution(words []string) []string {
	location := [26]int8{
		1, 2, 2, 1, 0, 1, 1, 1, 0, 1, 1, 1, 2, 2, 0, 0, 0, 0, 1, 0, 0, 2, 0, 2, 0, 2,
	}
	ans := make([]string, 0)
	for _, str := range words {
		i := 0
		idx := 0
		cur := int8(-1)
		for ; i < len(str); i++ {
			if str[i] >= 'a' && str[i] <= 'z' {
				idx = int(str[i] - 'a')
			}
			if str[i] >= 'A' && str[i] <= 'Z' {
				idx = int(str[i] - 'A')
			}
			if cur == -1 {
				cur = location[idx]
				continue
			}
			if location[idx] != cur {
				break
			}
		}
		if i == len(str) {
			ans = append(ans, str)
		}

	}
	return ans
}
