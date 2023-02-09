package Solution

func intersection2306(a, b map[string]struct{}) int64 {
	i := int64(0)
	for k := range a {
		if _, ok := b[k]; ok {
			i++
		}
	}
	return i
}
func Solution(ideas []string) int64 {
	byteWithSuffix := make([]map[string]struct{}, 26)
	for _, w := range ideas {
		if byteWithSuffix[w[0]-'a'] == nil {
			byteWithSuffix[w[0]-'a'] = make(map[string]struct{})
		}
		byteWithSuffix[w[0]-'a'][w[1:]] = struct{}{}
	}

	ans := int64(0)
	for b := 'b'; b <= 'z'; b++ {
		x := int64(len(byteWithSuffix[b-'a']))
		for pre := 'a'; pre < b; pre++ {
			y := int64(len(byteWithSuffix[pre-'a']))
			diff := intersection2306(byteWithSuffix[pre-'a'], byteWithSuffix[b-'a'])
			ans += (x - diff) * (y - diff)
		}
	}
	return ans * 2
}
