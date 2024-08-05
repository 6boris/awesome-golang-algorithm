package Solution

func Solution(arr []string, k int) string {
	count := make(map[string]int)
	for _, s := range arr {
		count[s]++
	}

	for i := range arr {
		if count[arr[i]] > 1 {
			continue
		}
		k--
		if k == 0 {
			return arr[i]
		}
	}
	return ""
}
