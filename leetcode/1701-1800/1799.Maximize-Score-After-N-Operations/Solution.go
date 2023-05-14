package Solution

import (
	"sort"
)

func gcd1799(a, b int) int {
	for a != b {
		if a > b {
			a, b = b, a-b
		} else {
			a, b = b, a
		}
	}
	return b
}

type gcdItem struct {
	i, j, gcd int
}

func findAns(gcdArray []gcdItem, used []bool, start, index, sum int, ans *int) {
	if index <= 0 || start >= len(gcdArray) {
		if sum > *ans {
			*ans = sum
		}
		return
	}
	findAns(gcdArray, used, start+1, index, sum, ans)
	x := gcdArray[start]
	if !(used[x.i] || used[x.j]) {
		used[x.i], used[x.j] = true, true
		findAns(gcdArray, used, start+1, index-1, sum+index*x.gcd, ans)
		used[x.i], used[x.j] = false, false
	}

}
func Solution(nums []int) int {
	ans := 0
	l := len(nums)
	used := make([]bool, l)
	gcdArray := make([]gcdItem, 0)
	for i := 0; i < l; i++ {
		for j := i + 1; j < l; j++ {
			x := gcd1799(nums[i], nums[j])
			gcdArray = append(gcdArray, gcdItem{i, j, x})
		}
	}
	sort.Slice(gcdArray, func(i, j int) bool {
		return gcdArray[i].gcd > gcdArray[j].gcd
	})
	index := l / 2
	findAns(gcdArray, used, 0, index, 0, &ans)
	return ans
}
