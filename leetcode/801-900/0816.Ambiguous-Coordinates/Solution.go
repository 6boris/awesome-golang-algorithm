package Solution

import "fmt"

func genParts(bs string) []string {
	if len(bs) == 1 {
		return []string{string(bs)}
	}

	l := len(bs)
	if bs[0] == '0' && bs[l-1] == '0' {
		return nil
	}
	if bs[0] == '0' {
		return []string{fmt.Sprintf("0.%s", bs[1:])}
	}
	if bs[l-1] == '0' {
		return []string{string(bs)}
	}
	ans := []string{string(bs)}
	// 1, 2 3 4
	for i := 1; i < l; i++ {
		ans = append(ans, fmt.Sprintf("%s.%s", bs[:i], bs[i:]))
	}
	return ans
}
func Solution(s string) []string {
	ll := len(s)
	bs := s[1 : ll-1]
	ans := make([]string, 0)
	keys := make(map[string]struct{})
	for i := 1; i < len(bs); i++ {
		l := genParts(bs[:i])
		r := genParts(bs[i:])
		if len(l) == 0 || len(r) == 0 {
			continue
		}
		for _, ls := range l {
			for _, rs := range r {
				k := fmt.Sprintf("(%s, %s)", ls, rs)
				if _, ok := keys[k]; !ok {
					ans = append(ans, k)
					keys[k] = struct{}{}
				}
			}
		}
	}
	return ans
}
