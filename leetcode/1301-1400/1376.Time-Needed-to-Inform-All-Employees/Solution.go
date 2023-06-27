package Solution

func Solution(n int, headID int, manager []int, informTime []int) int {
	managers := make(map[int][]int)
	for e, m := range manager {
		if _, ok := managers[m]; !ok {
			managers[m] = make([]int, 0)
		}
		managers[m] = append(managers[m], e)
	}
	var dfs func(int, int, *int)
	dfs = func(i, dis int, ans *int) {
		v, ok := managers[i]
		if !ok {
			if dis > *ans {
				*ans = dis
			}
			return
		}
		for _, employ := range v {
			dfs(employ, dis+informTime[i], ans)
		}
	}
	ans := 0
	dfs(headID, 0, &ans)
	return ans
}
