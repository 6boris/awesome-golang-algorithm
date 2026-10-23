package Solution

import "sort"

func parse1096(expr string) map[string]bool {
	var parts []string
	depth := 0
	lastIndex := 0
	for i := 0; i < len(expr); i++ {
		if expr[i] == '{' {
			depth++
		} else if expr[i] == '}' {
			depth--
		} else if expr[i] == ',' && depth == 0 {
			parts = append(parts, expr[lastIndex:i])
			lastIndex = i + 1
		}
	}

	if len(parts) > 0 {
		parts = append(parts, expr[lastIndex:])
		res := make(map[string]bool)
		for _, p := range parts {
			for k := range parse1096(p) {
				res[k] = true // 自动去重
			}
		}
		return res
	}

	var factors []string
	i := 0
	for i < len(expr) {
		if expr[i] == '{' {
			start := i
			d := 0
			for i < len(expr) {
				if expr[i] == '{' {
					d++
				} else if expr[i] == '}' {
					d--
				}
				i++
				if d == 0 {
					break
				}
			}
			factors = append(factors, expr[start:i])
		} else {
			factors = append(factors, string(expr[i]))
			i++
		}
	}

	if len(factors) == 1 {
		f := factors[0]
		if len(f) > 0 && f[0] == '{' && f[len(f)-1] == '}' {
			return parse1096(f[1 : len(f)-1])
		}
		return map[string]bool{f: true}
	}

	res := map[string]bool{"": true}
	for _, factor := range factors {
		subSet := parse1096(factor)
		nextRes := make(map[string]bool)
		for w1 := range res {
			for w2 := range subSet {
				nextRes[w1+w2] = true
			}
		}
		res = nextRes
	}

	return res
}

func Solution(expression string) []string {
	set := parse1096(expression)
	var result []string
	for word := range set {
		result = append(result, word)
	}
	sort.Strings(result)
	return result
}
