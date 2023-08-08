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

// Solution2
type operation struct {
	val int
	op  byte
}

func calResult(a, b int, op byte) int {
	if op == '*' {
		return a * b
	}
	if op == '+' {
		return a + b
	}
	return a - b
}

func diffWaysToCompute2(expression string) []int {
	opts := make([]operation, 0)
	base := 1
	c := 0
	// 245
	for idx := len(expression) - 1; idx >= 0; idx-- {
		if expression[idx] >= '0' && expression[idx] <= '9' {
			c = base*int(expression[idx]-'0') + c
			base *= 10
			continue
		}
		opts = append(opts, operation{val: c}, operation{op: expression[idx]})
		base, c = 1, 0
	}
	opts = append(opts, operation{val: c})

	length := len(opts)
	if length == 1 {
		return []int{opts[0].val}
	}
	for s, e := 0, length-1; s < e; s, e = s+1, e-1 {
		opts[s], opts[e] = opts[e], opts[s]
	}

	ans := make([]int, 0)
	dp := make([][]map[int]int, length)
	// 需要明确dp[i][j] 能计算出多少种结果
	for i := 0; i < length; i++ {
		dp[i] = make([]map[int]int, length)
		for j := 0; j < length; j++ {
			dp[i][j] = make(map[int]int)
		}
		if i&1 == 0 {
			dp[i][i][opts[i].val] = 1
		}
	}

	// 2 - 1 - 1-1
	// n=2 0-2, 2-4, 4-6
	// n=3,0-4, 2-6
	// n=4,0-6
	// n = 4 end=6
	// length=7
	// 2-1
	// 2 有一个操作数，3有两个，4有3个
	// 2 = 1+0
	// 3 = 2+1
	// 4 = 3 + 2
	// length - 1 - n+1 -n+2

	for n := 2; n <= length/2+1; n++ {
		for start := 0; start < length+2-2*n; start += 2 {
			end := start + 2*(n-1)
			for k := start; k < end; k += 2 {
				for v1, c1 := range dp[start][k] {
					for v2, c2 := range dp[k+2][end] {
						r := calResult(v1, v2, opts[k+1].op)
						dp[start][end][r] += c1 * c2
					}
				}
			}
		}
	}
	for idx := length - 1; idx > 0; idx -= 2 {
		op := opts[idx-1].op
		for v1, c1 := range dp[idx][length-1] {
			for v2, c2 := range dp[0][idx-2] {
				r := calResult(v2, v1, op)
				for loop := c1 * c2; loop > 0; loop-- {
					ans = append(ans, r)
				}
			}
		}
	}
	return ans
}
