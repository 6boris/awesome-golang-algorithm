package Solution

func merge_1(nums1 []int, m int, nums2 []int, n int) {
	sorted := make([]int, 0, m+n)
	p, q := 0, 0
	for {
		if p == m {
			sorted = append(sorted, nums2[q:]...)
			break
		}
		if q == n {
			sorted = append(sorted, nums1[p:]...)
			break
		}
		if nums1[p] < nums2[q] {
			sorted = append(sorted, nums1[p])
			p++
		} else {
			sorted = append(sorted, nums2[q])
			q++
		}
	}
	copy(nums1, sorted)
}

func merge_2(nums1 []int, m int, nums2 []int, n int) {
	p, q := m-1, n-1
	tail := m + n - 1
	for p >= 0 && q >= 0 {
		if nums1[p] > nums2[q] {
			nums1[tail] = nums1[p]
			p--
		} else {
			nums1[tail] = nums2[q]
			q--
		}
		tail--
	}
	if p == -1 && q >= 0 {
		for q >= 0 {
			nums1[tail] = nums2[q]
			q--
			tail--
		}
	}
	if q == -1 && p >= 0 {
		for p >= 0 {
			nums1[tail] = nums1[p]
			p--
			tail--
		}
	}
}

func merge_3(nums1 []int, m int, nums2 []int, n int) {
	p, q := m-1, n-1
	tail := m + n - 1
	for p >= 0 || q >= 0 {
		if p == -1 {
			nums1[tail] = nums2[q]
			q--
		} else if q == -1 {
			nums1[tail] = nums1[p]
			p--
		} else if nums1[p] > nums2[q] {
			nums1[tail] = nums1[p]
			p--
		} else {
			nums1[tail] = nums2[q]
			q--
		}
		tail--
	}
}
