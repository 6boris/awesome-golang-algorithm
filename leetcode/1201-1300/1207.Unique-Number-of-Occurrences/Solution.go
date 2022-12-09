package Solution

func Solution(arr []int) bool {
	// -1=1001, -1000=2000
	bucket := make([]int, 2001)
	for _, n := range arr {
		if n < 0 {
			n = ^n + 1 + 1000
		}
		bucket[n]++
	}
	occur := make([]bool, 1001)
	for _, c := range bucket {
		if c == 0 {
			continue
		}
		if occur[c] {
			return false
		}
		occur[c] = true
	}
	return true
}
