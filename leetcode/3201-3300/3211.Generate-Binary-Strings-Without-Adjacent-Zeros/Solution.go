package Solution

func Solution(n int) []string {
	var res []string
	var dfs func(int, string, bool)
	dfs = func(index int, cur string, isZero bool) {
		if index == n {
			res = append(res, cur)
			return
		}
		if !isZero {
			dfs(index+1, cur+"0", true)
		}
		dfs(index+1, cur+"1", false)
	}
	dfs(0, "", false)
	return res
}
