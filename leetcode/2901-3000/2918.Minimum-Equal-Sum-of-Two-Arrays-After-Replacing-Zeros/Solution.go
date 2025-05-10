package Solution

func Solution(nums1 []int, nums2 []int) int64 {
	a, b := 0, 0
	a0, b0 := 0, 0
	for _, n := range nums1 {
		a += n
		if n == 0 {
			a0++
		}
	}

	for _, n := range nums2 {
		b += n
		if n == 0 {
			b0++
		}
	}
	if a0 == 0 && b0 == 0 {
		if a != b {
			return -1
		}
		return int64(a)
	}

	if a0 != 0 && b0 != 0 {
		return max(int64(a+a0), int64(b+b0))
	}
	if a0 == 0 {
		// b0 != 0
		if b+b0 <= a {
			return int64(a)
		}
		return -1
	}
	if a+a0 <= b {
		return int64(b)
	}
	return -1
}
