package Solution

func Solution(s string, dictionary []string) int {
	dic := make(map[string]struct{})
	for _, str := range dictionary {
		dic[str] = struct{}{}
	}
	if _, ok := dic[s]; ok {
		return 0
	}

	cache := make(map[string]int)

	var dfs func(string) int
	dfs = func(str string) int {
		if _, ok := dic[str]; ok || str == "" {
			return 0
		}
		if v, ok := cache[str]; ok {
			return v
		}
		extra := -1
		for i := 1; i <= len(str); i++ {
			part := str[0:i]
			tmp := i // delete characters
			if _, ok := dic[part]; ok {
				tmp = 0
			}
			x := dfs(str[i:])
			if extra == -1 || x+tmp < extra {
				extra = x + tmp
			}

		}
		cache[str] = extra
		return extra
	}
	return dfs(s)
}
