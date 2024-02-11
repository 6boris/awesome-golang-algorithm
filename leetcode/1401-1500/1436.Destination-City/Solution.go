package Solution

func Solution(paths [][]string) string {
	in := make(map[string]struct{})
	for _, p := range paths {
		in[p[0]] = struct{}{}
	}
	for _, p := range paths {
		if _, ok := in[p[1]]; !ok {
			return p[1]
		}
	}
	return ""
}
