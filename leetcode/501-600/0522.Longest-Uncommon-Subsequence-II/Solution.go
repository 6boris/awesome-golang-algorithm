package Solution

func sub522(source string, l, cur int, exist map[string]struct{}, str []byte, result *[]string) {
	if l == 0 {
		y := string(str)
		if _, ok := exist[y]; !ok {
			*result = append(*result, y)
			exist[y] = struct{}{}
		}
		return
	}
	if cur >= len(source) {
		return
	}
	str = append(str, source[cur])
	sub522(source, l-1, cur+1, exist, str, result)
	str = str[:len(str)-1]
	sub522(source, l, cur+1, exist, str, result)
}

type lus struct {
	source  string
	subsets []string
	exist   map[string]struct{}
}

func Solution(strs []string) int {
	length := len(strs)
	ls := make([]lus, 0)
	exist := make(map[string]int)
	for i := range strs {
		exist[strs[i]]++
		if exist[strs[i]] > 1 {
			continue
		}
		ls = append(ls, lus{source: strs[i], subsets: make([]string, 0), exist: make(map[string]struct{})})
		li := len(ls) - 1
		for l := len(strs[i]); l > 0; l-- {
			sub522(strs[i], l, 0, ls[li].exist, []byte{}, &ls[li].subsets)
		}
	}
	length = len(ls)
	ans := -1
	for i := 0; i < length; i++ {
		if exist[ls[i].source] > 1 {
			continue
		}
		for _, sub := range ls[i].subsets {
			j := 0
			for ; j < length; j++ {
				if j == i {
					continue
				}
				if _, ok := ls[j].exist[sub]; ok {
					break
				}
			}
			if j == length && len(sub) > ans {
				ans = len(sub)
				break
			}
		}

	}
	return ans
}
