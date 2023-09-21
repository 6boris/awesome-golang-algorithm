package Solution

func Min(a int, b int) int {
	if a > b {
		return b
	}
	return a
}

func Max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	m, n := len(nums1), len(nums2)
	halfLen := (m + n + 1) / 2

	if m == 0 || n == 0 {
		nums := append(nums1, nums2...)
		if len(nums)%2 == 1 {
			return float64(nums[halfLen-1])
		} else {
			return (float64(nums[halfLen-1]) + float64(nums[halfLen])) / 2.0
		}
	}

	// make sure m < n, so j = halfLen - i is always greater than zero
	var A, B []int
	if m < n {
		A, B = nums1, nums2
	} else {
		m, n = n, m
		A, B = nums2, nums1
	}

	// find i in [0, m]
	// especial, if i == 0 means A is the right part, if i == m means A is the left part
	iMin, iMax := 0, m
	for iMin <= iMax {
		i := (iMin + iMax) / 2
		j := halfLen - i

		if i > 0 && j >= 0 && j < n && A[i-1] > B[j] {
			// Ai > Bj+1, i need be smaller
			iMax = i - 1
		} else if j > 0 && i >= 0 && i < m && B[j-1] > A[i] {
			// Bj > Ai+1, i need be greater
			iMin = i + 1
		} else {
			var leftPartMax, rightPartMin float64
			if i == 0 {
				leftPartMax = float64(B[j-1])
			} else if j == 0 {
				leftPartMax = float64(A[i-1])
			} else {
				leftPartMax = float64(Max(A[i-1], B[j-1]))
			}

			if (m+n)%2 == 1 {
				// m + n is odd
				return leftPartMax
			}

			if i == m {
				rightPartMin = float64(B[j])
			} else if j == n {
				rightPartMin = float64(A[i])
			} else {
				rightPartMin = float64(Min(A[i], B[j]))
			}

			// m + n is even
			return (leftPartMax + rightPartMin) / 2.0
		}
	}

	return -1.0
}

func findMedianSortedArrays1(nums1 []int, nums2 []int) float64 {
	l1 := len(nums1)
	l2 := len(nums2)
	targetIndex := (l1 + l2) / 2
	index := 0
	i, j := 0, 0
	a, b := 0, 0
	for i < l1 && j < l2 && index <= targetIndex {
		index++
		b = a
		if nums1[i] < nums2[j] {
			a = nums1[i]
			i++
			continue
		}
		a = nums2[j]
		j++
	}
	for ; i < l1 && index <= targetIndex; i++ {
		b = a
		a = nums1[i]
		index++
	}
	for ; j < l2 && index <= targetIndex; j++ {
		b = a
		a = nums2[j]
		index++
	}
	if (l1+l2)&1 == 1 {
		return float64(a)
	}
	return float64(a+b) / 2
}
