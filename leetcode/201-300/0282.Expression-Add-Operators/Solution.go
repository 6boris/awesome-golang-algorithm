package Solution

func Solution(num string, target int) []string {
	ans := make([]string, 0)
	in := make(map[string]struct{})
	var dfs func(int, string)
	dfs = func(start int, cur string) {
		if start >= len(num) {
			r := cur[:len(cur)-1]
			if _, ok := in[r]; !ok {
				ans = append(ans, r)
			}
			in[r] = struct{}{}
			return
		}
		if num[start] == '0' {
			cur += "0"
			dfs(start+1, cur+"+")
			dfs(start+1, cur+"-")
			dfs(start+1, cur+"*")
			return
		}

		for end := start; end < len(num); end++ {
			now := cur + num[start:end+1]
			dfs(end+1, now+"+")
			dfs(end+1, now+"-")
			dfs(end+1, now+"*")
		}
	}
	dfs(0, "")
	var cal func(string) int
	result := make([]string, 0)
	cal = func(str string) int {
		n := 0
		nstack := make([]int, 0)
		cstack := make([]byte, 0)
		for i := range str {
			if str[i] >= '0' && str[i] <= '9' {
				n = n*10 + int(str[i]-'0')
				continue
			}
			nstack = append(nstack, n)
			n = 0
			if str[i] != '*' {
				for len(cstack) > 0 && len(nstack) >= 2 {
					l := len(nstack)
					a, b := nstack[l-2], nstack[l-1]
					nstack = nstack[:l-2]
					top := cstack[len(cstack)-1]
					cstack = cstack[:len(cstack)-1]
					if top == '*' {
						nstack = append(nstack, a*b)
					}
					if top == '+' {
						nstack = append(nstack, a+b)
					}
					if top == '-' {
						nstack = append(nstack, a-b)
					}
				}
			}
			cstack = append(cstack, str[i])
		}
		nstack = append(nstack, n)

		for len(cstack) > 0 && len(nstack) >= 2 {
			l := len(nstack)
			a, b := nstack[l-2], nstack[l-1]
			nstack = nstack[:l-2]
			top := cstack[len(cstack)-1]
			cstack = cstack[:len(cstack)-1]
			if top == '*' {
				nstack = append(nstack, a*b)
			}
			if top == '+' {
				nstack = append(nstack, a+b)
			}
			if top == '-' {
				nstack = append(nstack, a-b)
			}
		}
		return nstack[0]
	}
	for _, s := range ans {
		if r := cal(s); r == target {
			result = append(result, s)
		}
	}
	return result
}
