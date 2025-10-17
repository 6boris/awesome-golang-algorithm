package Solution

func Solution(s string, k int) int {
	n := len(s)
	left := make([][3]int, n)
	right := make([][3]int, n)

	num, mask, count := 0, 0, 0
	for i := 0; i < n-1; i++ {
		binary := 1 << (s[i] - 'a')
		if mask&binary == 0 {
			count++
			if count <= k {
				mask |= binary
			} else {
				num++
				mask = binary
				count = 1
			}
		}
		left[i+1][0] = num
		left[i+1][1] = mask
		left[i+1][2] = count
	}

	num, mask, count = 0, 0, 0
	for i := n - 1; i > 0; i-- {
		binary := 1 << (s[i] - 'a')
		if mask&binary == 0 {
			count++
			if count <= k {
				mask |= binary
			} else {
				num++
				mask = binary
				count = 1
			}
		}
		right[i-1][0] = num
		right[i-1][1] = mask
		right[i-1][2] = count
	}

	maxVal := 0
	for i := 0; i < n; i++ {
		seg := left[i][0] + right[i][0] + 2
		totMask := left[i][1] | right[i][1]
		totCount := 0
		for totMask != 0 {
			totMask = totMask & (totMask - 1)
			totCount++
		}
		if left[i][2] == k && right[i][2] == k && totCount < 26 {
			seg++
		} else if min(totCount+1, 26) <= k {
			seg--
		}
		if seg > maxVal {
			maxVal = seg
		}
	}
	return maxVal
}
