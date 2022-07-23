package Solution

func Solution(target, arr []int) bool {
	check := make([]int, 1001)
	for idx := 0; idx < len(target); idx++ {
		check[target[idx]]++
		check[arr[idx]]--
	}
	for idx := 0; idx < 1001; idx++ {
		if check[idx] != 0 {
			return false
		}
	}
	return true
}
