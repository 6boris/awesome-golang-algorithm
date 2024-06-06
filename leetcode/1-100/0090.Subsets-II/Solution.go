package Solution

func Solution(nums []int) [][]int {
	ans := [][]int{{}}
	cache := make(map[[21]int]struct{})

	var (
		dfs   func(int, int, [21]int)
		toRes func([21]int) []int
	)

	toRes = func(path [21]int) []int {
		r := make([]int, 0)
		for i := 0; i <= 10; i++ {
			if path[i] == 0 {
				continue
			}
			for count := path[i]; count > 0; count-- {
				r = append(r, i)
			}
		}
		for i := 11; i < 21; i++ {
			if path[i] == 0 {
				continue
			}
			for count := path[i]; count > 0; count-- {
				r = append(r, 10-i)
			}
		}
		return r
	}

	dfs = func(index, l int, path [21]int) {
		if l == 0 {
			if _, ok := cache[path]; !ok {
				ans = append(ans, toRes(path))
				cache[path] = struct{}{}
			}
			return
		}
		if index == len(nums) {
			return
		}
		cur := nums[index]
		if cur < 0 {
			cur = 10 - cur
		}
		path[cur]++
		dfs(index+1, l-1, path)
		path[cur]--
		dfs(index+1, l, path)
	}
	for l := 1; l <= len(nums); l++ {
		dfs(0, l, [21]int{})
	}

	return ans
}
