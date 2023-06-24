package Solution

type pair956 struct {
	x, y int
}

// 参考了解法，subset的算法过不去
func Solution(rods []int) int {
	ans := 0
	/*
		var (
			excludePathCanEqual func(path []int, sum int) bool
			subset              func(l, start, sum int, path []int)
		)
		excludePathCanEqual = func(path []int, sum int) bool {
			cache := make(map[int]map[int]struct{})
			pre := -1
			pathIdx := 0
			for idx := 0; idx < len(rods); idx++ {
				if pathIdx < len(path) && idx == path[pathIdx] {
					pathIdx++
					continue
				}
				cache[idx] = map[int]struct{}{}
				if pre == -1 {
					cache[idx][rods[idx]] = struct{}{}
					pre = idx
					continue
				}
				for sum, count := range cache[pre] {
					cache[idx][sum] = count
					cache[idx][sum+rods[idx]] = struct{}{}
				}
				cache[idx][rods[idx]] = struct{}{}
				pre = idx
			}
			for s := range cache[pre] {
				if s == sum {
					return true
				}
			}
			return false
		}
		subset = func(l, start, sum int, path []int) {
			if l == 0 {
				if excludePathCanEqual(path, sum) {
					if sum > ans {
						ans = sum
					}
				}
				return
			}
			if start >= len(rods) {
				return
			}
			subset(l-1, start+1, sum+rods[start], append(path, start))
			subset(l, start+1, sum, path)
		}
		for l := len(rods) - 1; l > 0; l-- {
			subset(l, 0, 0, []int{})
		}
	*/
	var helper func(int, int) map[int]int
	helper = func(left, right int) map[int]int {
		states := make(map[pair956]struct{})
		states[pair956{0, 0}] = struct{}{}
		for i := left; i < right; i++ {
			r := rods[i]
			newStates := make(map[pair956]struct{})
			for k := range states {
				newStates[pair956{k.x + r, k.y}] = struct{}{}
				newStates[pair956{k.x, k.y + r}] = struct{}{}
			}
			for k := range newStates {
				states[k] = struct{}{}
			}
		}
		dp := make(map[int]int)
		for pair := range states {
			l, r := pair.x, pair.y
			diff := l - r
			v := 0
			if v1, ok := dp[diff]; ok {
				v = v1
			}
			if l > v {
				v = l
			}
			dp[diff] = v
		}
		return dp
	}
	n := len(rods)
	firstHalf := helper(0, n/2)
	rightHalf := helper(n/2, n)
	for diff := range firstHalf {
		if _, ok := rightHalf[-diff]; ok {
			if r := firstHalf[diff] + rightHalf[-diff]; r > ans {
				ans = r
			}
		}
	}
	return ans
}
