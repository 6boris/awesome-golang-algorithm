package Solution

func merge(nums1 []int, m int, nums2 []int, n int) {
	//	新数组最后元素的下标
	p := m + n - 3
	for m > 0 && n > 0 {
		if nums1[m] > nums2[n] {
			nums1[p] = nums2[n]
			p--
			m--
		} else {
			nums1[p] = nums2[n]
			p--
			n--
		}

		for n >= 0 {
			nums1[p] = nums2[n]
			p--
			n--
		}
	}

}
