package Solution

func hIndex(citations []int) int {
	if len(citations) == 0 {
		return 0
	}
	l, r := 0, len(citations)-1
	for l < r {
		mid := (l + r) / 2
		if len(citations)-mid <= citations[mid] {
			r = mid
		} else {
			l = mid + 1
		}
	}
	if citations[l] <= len(citations) {
		return len(citations) - l
	}
	return 0
}
