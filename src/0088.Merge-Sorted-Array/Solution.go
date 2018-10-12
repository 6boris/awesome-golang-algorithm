package Solution

func merge(nums1 []int, m int, nums2 []int, n int) {
	//	新数组最后元素的下标
	p := m + n - 1
	m--
	n--
	for m >= 0 && n >= 0 {
		if nums1[m] > nums2[n] {
			nums1[p] = nums1[m]
			p--
			m--
		} else {
			nums1[p] = nums2[n]
			p--
			n--
		}
	}

	for n >= 0 {
		nums1[p] = nums2[n]
		p--
		n--
	}
}
