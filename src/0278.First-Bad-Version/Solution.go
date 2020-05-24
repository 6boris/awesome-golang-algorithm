package Solution

//	直接二分
func firstBadVersion(n int) int {
	left, right := 1, n

	for left < right {
		mid := (right-left)/2 + left
		if isBadVersion(mid) {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}

func isBadVersion(version int) bool {
	m := make(map[int]bool)
	m[3] = false
	m[4] = true
	m[5] = false

	if v, ok := m[version]; ok {
		return v
	}

	return false
}
