package Solution

func Solution(logs [][]int) int {
	bucket := make([]int, 101)
	for _, log := range logs {
		for year := log[0]; year < log[1]; year++ {
			bucket[year-1950]++
		}
	}
	targetIdx := 0
	for idx := 1; idx < 101; idx++ {
		if bucket[idx] > bucket[targetIdx] {
			targetIdx = idx
		}
	}
	return targetIdx + 1950
}
