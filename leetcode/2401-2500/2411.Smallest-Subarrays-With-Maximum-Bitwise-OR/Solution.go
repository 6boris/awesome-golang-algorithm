package Solution

import "sort"

func oneArray(n int) [32]int {
	res := [32]int{}
	index := 31
	for ; n > 0; index-- {
		if n&1 == 1 {
			res[index]++
		}
		n >>= 1
	}
	return res
}

func Solution(nums []int) []int {
	l := len(nums)
	ans := make([]int, l)
	ans[l-1] = 1

	or := make([]int, l)
	or[l-1] = nums[l-1]
	for i := l - 2; i >= 0; i-- {
		or[i] = nums[i] | or[i+1]
	}

	oneCount := make([][32]int, l)
	oneCount[0] = oneArray(nums[0])
	for i := 1; i < l; i++ {
		cur := oneArray(nums[i])
		for j := 0; j < 32; j++ {
			oneCount[i][j] = oneCount[i-1][j] + cur[j]
		}
	}
	var ok func(int, int, int) bool
	ok = func(start, end, target int) bool {
		tmp := oneCount[end]
		if start != 0 {
			for i := 0; i < 32; i++ {
				tmp[i] -= oneCount[start-1][i]
			}
		}
		base := 1
		num := 0
		for i := 31; i >= 0; i-- {
			x := 1
			if tmp[i] == 0 {
				x = 0
			}
			num += x * base
			base <<= 1
		}
		return num == target
	}
	for i := l - 2; i >= 0; i-- {
		// i = 3
		index := sort.Search(l-i, func(j int) bool {
			return ok(i, i+j, or[i])
		})
		ans[i] = index + 1
	}
	return ans
}
