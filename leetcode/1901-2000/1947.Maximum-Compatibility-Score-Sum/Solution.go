package Solution

func Solution(students [][]int, mentors [][]int) int {
	var (
		dfs   func(int, []bool) int
		match func(int, int) int
	)
	match = func(i, j int) int {
		ans := 0
		for k := range students[i] {
			if mentors[j][k] == students[i][k] {
				ans++
			}
		}
		return ans
	}

	dfs = func(index int, used []bool) int {
		cur := 0
		if index == len(students) {
			return cur
		}

		for i := range mentors {
			if used[i] {
				continue
			}
			used[i] = true
			cur = max(cur, match(index, i)+dfs(index+1, used))
			used[i] = false
		}
		return cur
	}
	return dfs(0, make([]bool, len(students)))
}
