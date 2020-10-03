package Solution

func generateParenthesis(n int) []string {
	return dp("", 0, n*2)
}

func dp(prefix string, depth int, n int) []string {
	if depth == 0 && n == 0 {
		return []string{prefix}
	}
	var output []string
	if depth < n {
		output = append(output, dp(prefix+"(", depth+1, n-1)...)
	}
	if depth > 0 {
		output = append(output, dp(prefix+")", depth-1, n-1)...)
	}
	return output
}
