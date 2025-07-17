package Solution

func longestAlternatingSubsequnece(l1, l2 []int) int {
	a, b := 0, 0
	ai, bi := 0, 0
	for ai < len(l1) && bi < len(l2) {
		if l1[ai] < l2[bi] {
			ai++
			a = b + 1
			continue
		}
		b = a + 1
		bi++
	}
	return max(a, b) + 1
}
func Solution(nums []int, k int) int {
	group := make([][]int, k)
	for i := range nums {
		nums[i] %= k
		group[nums[i]] = append(group[nums[i]], i)
	}
	// 1,0,1,0,1
	ans, l := 0, len(nums)
	sums := make(map[int]struct{})
	for i := 0; i < l; i++ {
		for j := i + 1; j < l; j++ {
			sums[nums[i]+nums[j]] = struct{}{}
		}
	}

	for sum := range sums {
		used := map[int]struct{}{}
		for i := 0; i < l; i++ {
			if _, ok := used[nums[i]]; ok {
				continue
			}
			used[nums[i]] = struct{}{}
			target := (k + sum - nums[i]) % k
			used[target] = struct{}{}

			ll := len(group[target])
			if ll == 0 {
				continue
			}
			if nums[i] == target {
				ans = max(ans, ll)
				continue
			}
			ans = max(ans, longestAlternatingSubsequnece(group[nums[i]], group[target]))
		}
	}
	return ans
}
