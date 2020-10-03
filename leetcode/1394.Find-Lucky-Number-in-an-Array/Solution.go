package Solution

func Solution(arr []int) int {
	n := len(arr)
	if n == 1 {
		if arr[0] == 1 {
			return 1
		} else {
			return -1
		}
	}
	hash := make(map[int]int)
	for _, val := range arr {
		hash[val]++
	}
	m := -1
	for _, val := range arr {
		if hash[val] == val && val > m {
			m = val
		}
	}
	return m
}
