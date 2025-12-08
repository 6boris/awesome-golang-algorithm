package Solution

func Solution(n int) int {
	cache := make(map[int]struct{})
	for i := 1; i <= n; i++ {
		cache[i*i] = struct{}{}
	}

	var ret int
	for i := 1; i <= n-1; i++ {
		for j := i + 1; j <= n; j++ {
			if _, ok := cache[i*i+j*j]; ok {
				ret += 2
			}
		}
	}
	return ret
}
