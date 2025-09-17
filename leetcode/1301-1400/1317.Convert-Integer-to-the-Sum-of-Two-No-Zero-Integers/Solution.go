package Solution

func ok(n int) bool {
	for n > 0 {
		mod := n % 10
		if mod == 0 {
			return false
		}
		n /= 10
	}
	return true
}

func Solution(n int) []int {
	nums := make(map[int]struct{})
	for i := 1; i <= 10000; i++ {
		if ok(i) {
			nums[i] = struct{}{}
		}
	}
	var diff int
	for k := range nums {
		diff = n - k
		if _, ok := nums[diff]; ok {
			return []int{k, diff}
		}
	}
	return []int{}
}
