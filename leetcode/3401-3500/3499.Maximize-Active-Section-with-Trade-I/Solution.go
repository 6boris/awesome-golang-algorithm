package Solution

func Solution(s string) int {
	n := len(s)
	cnt1 := 0
	for _, c := range s {
		if c == '1' {
			cnt1++
		}
	}

	zeroBlocks := []int{}
	i := 0
	for i < n {
		start := i
		for i < n && s[i] == s[start] {
			i++
		}
		if s[start] == '0' {
			zeroBlocks = append(zeroBlocks, i-start)
		}
	}

	m := len(zeroBlocks)
	if m < 2 {
		return cnt1
	}

	bestGain := 0 // Optimal Increment
	for j := 0; j < m-1; j++ {
		bestGain = max(bestGain, zeroBlocks[j]+zeroBlocks[j+1])
	}

	return cnt1 + bestGain
}
