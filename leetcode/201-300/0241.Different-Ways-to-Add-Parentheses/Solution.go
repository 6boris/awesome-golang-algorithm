package Solution

import "strconv"

func diffWaysToCompute(input string) []int {
	return ways(input, map[string][]int{})
}

func calculate(a, b int, operate rune) int {
	switch operate {
	case '+':
		return a + b
	case '-':
		return a - b
	case '*':
		return a * b
	default:
		panic("operate not exist")
	}
}

func ways(input string, cache map[string][]int) []int {
	if _, ok := cache[input]; ok {
		return cache[input]
	}

	ans := []int{}

	for i := 0; i < len(input); i++ {
		ch := input[i]

		if ch == '+' || ch == '-' || ch == '*' {
			left := input[:i]
			right := input[i+1:]

			l := ways(left, cache)
			r := ways(right, cache)

			for _, a := range l {
				for _, b := range r {
					ans = append(ans, calculate(a, b, rune(ch)))
				}
			}
		}
	}

	if len(ans) == 0 {
		number, _ := strconv.Atoi(input)
		ans = append(ans, number)
	}

	cache[input] = ans

	return ans
}
