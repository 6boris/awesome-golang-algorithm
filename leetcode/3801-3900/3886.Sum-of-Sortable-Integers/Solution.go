package Solution

import "sort"

func Solution(nums []int) int {
	n := len(nums)
	var fac []int

	for x := 1; x*x <= n; x++ {
		if n%x == 0 {
			fac = append(fac, x)
			if n/x != x {
				fac = append(fac, n/x)
			}
		}
	}

	target := make([]int, n)
	copy(target, nums)
	sort.Ints(target)

	kmp := func(nums1, nums2 []int) bool {
		n, m := len(nums1), len(nums2)
		pi := make([]int, m)
		j := 0

		for i := 1; i < m; i++ {
			for j > 0 && nums2[i] != nums2[j] {
				j = pi[j-1]
			}
			if nums2[i] == nums2[j] {
				j++
			}
			pi[i] = j
		}

		j = 0
		for i := 0; i < n; i++ {
			for j > 0 && nums1[i] != nums2[j] {
				j = pi[j-1]
			}
			if nums1[i] == nums2[j] {
				j++
			}
			if j == m {
				return true
			}
		}
		return false
	}

	check := func(nums1, nums2 []int) bool {
		// Check whether cyclic rotations of nums1 can produce nums2
		doubled := make([]int, len(nums1)*2)
		copy(doubled, nums1)
		copy(doubled[len(nums1):], nums1)
		return kmp(doubled, nums2)
	}

	ans := 0

	for _, d := range fac {
		allGood := true
		for left := 0; left < n; left += d {
			right := left + d - 1

			if !check(nums[left:right+1], target[left:right+1]) {
				allGood = false
				break
			}
		}
		if allGood {
			ans += d
		}
	}

	return ans
}
