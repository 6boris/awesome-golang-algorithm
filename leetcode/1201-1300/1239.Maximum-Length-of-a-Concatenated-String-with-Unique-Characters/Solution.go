package Solution

func Solution(arr []string) int {
	filter := make([]string, 0)
	for _, x := range arr {
		tmp := [26]bool{}
		add := true
		for _, b := range []byte(x) {
			if tmp[b-'a'] {
				add = false
				break
			}
			tmp[b-'a'] = true
		}
		if add {
			filter = append(filter, x)
		}
	}
	if len(filter) == 0 {
		return 0
	}

	ans := 0
	var (
		subset    func(int, int, int, [26]bool)
		canSelect func(a [26]bool, b string) bool
	)
	canSelect = func(a [26]bool, b string) bool {
		for _, v := range b {
			if a[v-'a'] {
				return false
			}
		}
		return true
	}
	subset = func(index, subsetLen, strLen int, path [26]bool) {
		if subsetLen == 0 {
			if strLen > ans {
				ans = strLen
			}
			return
		}
		if index >= len(filter) {
			return
		}
		if canSelect(path, filter[index]) {
			for _, v := range filter[index] {
				path[v-'a'] = true
			}
			subset(index+1, subsetLen-1, strLen+len(filter[index]), path)
			for _, v := range filter[index] {
				path[v-'a'] = false
			}
		}
		subset(index+1, subsetLen, strLen, path)
	}
	for l := 1; l <= len(filter); l++ {
		subset(0, l, 0, [26]bool{})
	}
	return ans
}
