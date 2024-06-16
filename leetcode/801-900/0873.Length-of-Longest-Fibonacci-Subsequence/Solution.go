package Solution

func Solution(arr []int) int {
	nums := make(map[int]struct{})
	for _, n := range arr {
		nums[n] = struct{}{}
	}
	var try func(int, int, *int)
	try = func(a, b int, c *int) {
		_, ok := nums[a+b]
		if !ok {
			return
		}
		*c++
		try(b, a+b, c)
	}
	ans := 0
	for i := 0; i < len(arr)-1; i++ {
		for j := i + 1; j < len(arr); j++ {
			c := 0
			try(arr[i], arr[j], &c)
			if c != 0 {
				ans = max(ans, c+2)
			}
		}
	}
	return ans
}
