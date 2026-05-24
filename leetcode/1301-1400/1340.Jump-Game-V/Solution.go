package Solution

func Solution(arr []int, d int) int {
	size := len(arr)
	cache := make(map[int]int)

	var dfs func(int) int
	dfs = func(index int) int {
		if v, ok := cache[index]; ok {
			return v
		}
		cur := 0
		for step := 1; step <= d && index-step >= 0; step++ {
			if arr[index-step] >= arr[index] {
				break
			}
			cur = max(cur, dfs(index-step))
		}
		for step := 1; step <= d && index+step < size; step++ {
			if arr[index+step] >= arr[index] {
				break
			}
			cur = max(cur, dfs(index+step))
		}
		cache[index] = cur + 1
		return cache[index]
	}

	var ret int
	for i := 0; i < size; i++ {
		ret = max(ret, dfs(i))
	}
	return ret
}
