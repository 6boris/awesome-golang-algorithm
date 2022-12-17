package Solution

import "strconv"

func Solution(tokens []string) int {
	stack := make([]int, len(tokens))
	idx := 0
	for _, n := range tokens {
		if !(n == "+" || n == "-" || n == "*" || n == "/") {
			nn, _ := strconv.Atoi(n)
			stack[idx] = nn
			idx++
			continue
		}
		a, b := stack[idx-2], stack[idx-1]
		idx--
		var c int
		if n == "+" {
			c = a + b
		}
		if n == "*" {
			c = a * b
		}
		if n == "/" {
			c = a / b
		}
		if n == "-" {
			c = a - b
		}

		stack[idx-1] = c
	}
	return stack[0]
}
