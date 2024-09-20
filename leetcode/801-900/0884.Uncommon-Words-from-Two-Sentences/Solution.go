package Solution

import "strings"

func Solution(s1 string, s2 string) []string {
	ans := make([]string, 0)
	ls1 := strings.Split(s1, " ")
	ls2 := strings.Split(s2, " ")
	in1, in2 := make(map[string]int), make(map[string]int)
	for _, w := range ls1 {
		in1[w]++
	}
	for _, w := range ls2 {
		in2[w]++
	}
	for w, c := range in1 {
		if c != 1 {
			continue
		}
		if _, ok := in2[w]; !ok {
			ans = append(ans, w)
		}
	}
	for w, c := range in2 {
		if c != 1 {
			continue
		}
		if _, ok := in1[w]; !ok {
			ans = append(ans, w)
		}
	}
	return ans
}
