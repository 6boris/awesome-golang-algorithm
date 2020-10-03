package Solution

//	动态规划
func longestValidParentheses(s string) int {
	maxans := 0
	dp := make([]int, len(s))

	for i := 1; i < len(s); i++ {
		if s[i] == ')' {
			if s[i-1] == '(' {
				if i >= 2 {
					dp[i] = dp[i-2] + 2
				} else {
					dp[i] = 2
				}
			} else if i-dp[i-1] > 0 && s[i-dp[i-1]-1] == '(' {
				if i-dp[i-1] >= 2 {
					dp[i] = dp[i-1] + dp[i-dp[i-1]-2] + 2
				} else {
					dp[i] = dp[i-1] + 2
				}
			}
		}
		maxans = Max(maxans, dp[i])
	}
	return maxans
}

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

//	Stack
func longestValidParentheses2(s string) int {
	maxans := 0
	stk := Stack{}

	stk.Push(-1)

	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			stk.Push(i)
		} else {
			stk.Pop()
			if stk.IsEmpty() {
				stk.Push(i)
			} else {
				maxans = Max(maxans, i-stk.Top().(int))
			}
		}
	}
	return maxans
}
