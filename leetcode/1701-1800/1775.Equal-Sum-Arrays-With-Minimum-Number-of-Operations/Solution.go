package Solution

func Solution(nums1 []int, nums2 []int) int {
	n1, n2 := len(nums1), len(nums2)
	max1, min1 := n1*6, n1
	max2, min2 := n2*6, n2
	if max1 < min2 || max2 < min1 {
		return -1
	}
	sum1, sum2 := 0, 0
	for i := range n1 {
		sum1 += nums1[i]
	}
	for i := range n2 {
		sum2 += nums2[i]
	}

	if sum1 > sum2 {
		nums1, nums2 = nums2, nums1
		sum1, sum2 = sum2, sum1
	}

	diff := sum2 - sum1
	if diff == 0 {
		return 0
	}
	cnt := [6]int{}

	for _, x := range nums1 {
		cnt[6-x]++
	}
	for _, y := range nums2 {
		cnt[y-1]++
	}

	ops := 0
	for i := 5; i > 0; i-- {
		if cnt[i] == 0 {
			continue
		}

		if i*cnt[i] < diff {
			diff -= i * cnt[i]
			ops += cnt[i]
		} else {
			ops += (diff + i - 1) / i
			diff = 0
			break
		}
	}

	if diff == 0 {
		return ops
	}
	return -1
}
