package Solution

import "sort"

func validateCode(str string) bool {
	for _, b := range str {
		if !(b >= 'a' && b <= 'z' || b >= 'A' && b <= 'Z' || b >= '0' && b <= '9' || b == '_') {
			return false
		}
	}
	return len(str) > 0
}

func Solution(code []string, businessLine []string, isActive []bool) []string {

	indies := make([]int, 0)
	order := map[string]int{
		"electronics": 1, "grocery": 2, "pharmacy": 3, "restaurant": 4,
	}
	for index := range code {
		if !validateCode(code[index]) || !isActive[index] {
			continue
		}
		if _, ok := order[businessLine[index]]; !ok {
			continue
		}
		indies = append(indies, index)
	}

	sort.Slice(indies, func(i, j int) bool {
		a, b := businessLine[indies[i]], businessLine[indies[j]]
		if a == b {
			return code[indies[i]] < code[indies[j]]
		}
		return order[a] < order[b]
	})

	ret := make([]string, len(indies))
	for i := range indies {
		ret[i] = code[indies[i]]
	}

	return ret
}
