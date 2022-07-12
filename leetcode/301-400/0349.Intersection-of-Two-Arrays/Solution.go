package Solution

import "sort"

func intersection_1(nums1 []int, nums2 []int) []int {
	m1, m2 := make(map[int]int), make(map[int]int)
	ans := make([]int, 0)
	for idx, val := range nums1 {
		m1[val] = idx
	}
	for idx, val := range nums2 {
		if _, ok := m1[val]; ok {
			m2[val] = idx
		}
	}
	for key := range m2 {
		ans = append(ans, key)
	}

	return ans
}

func intersection_2(nums1 []int, nums2 []int) []int {
	sort.Ints(nums1)
	sort.Ints(nums2)
	ans := make([]int, 0)
	for i, j := 0, 0; i < len(nums1) && j < len(nums2); {
		x, y := nums1[i], nums2[j]
		if x == y && (len(ans) == 0 || ans[len(ans)-1] != x) {
			ans = append(ans, x)
			i++
			j++
		} else if x > y {
			j++
		} else {
			i++
		}
	}
	return ans
}

func intersection_3(nums1, nums2 []int) []int {
	ans := make([]int, 0)
	buckets := make([]int, 1001)
	for _, n := range nums1 {
		buckets[n] = 1
	}
	for _, n := range nums2 {
		if buckets[n] == 1 {
			ans = append(ans, n)
			buckets[n]++
		}
	}
	return ans
}
