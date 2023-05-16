package Solution

import (
	"sort"
)

func getSubsets(length, sum, start int, toppingCosts []int, next *[]int, in map[int]struct{}) {
	if length == 0 {
		if _, ok := in[sum]; !ok {
			*next = append(*next, sum)
			in[sum] = struct{}{}
		}
		return
	}
	if start >= len(toppingCosts) {
		return
	}
	getSubsets(length, sum, start+1, toppingCosts, next, in)
	getSubsets(length-1, sum+toppingCosts[start], start+1, toppingCosts, next, in)
	getSubsets(length-1, sum+toppingCosts[start]*2, start+1, toppingCosts, next, in)
}

// 二分搜索? 将topping里面的数据都计算出来。然后开始判断
func Solution(baseCosts []int, toppingCosts []int, target int) int {
	next := make([]int, 0)
	in := make(map[int]struct{})
	for l := 1; l <= len(toppingCosts); l++ {
		getSubsets(l, 0, 0, toppingCosts, &next, in)
	}
	sort.Ints(next)
	ans := -1
	for _, se := range baseCosts {
		diff := se - target
		if diff == 0 {
			return target
		}

		if diff > 0 {
			if ans == -1 {
				ans = se
				continue
			}

			d1 := target - ans
			if d1 < 0 {
				d1 = -d1
			}
			if diff < d1 || (diff == d1 && ans > se) {
				ans = se
			}
			continue
		}

		searchTarget := -diff
		if ans == -1 {
			ans = se
		} else {
			d1 := target - ans
			if d1 < 0 {
				d1 = -d1
			}
			if d1 > searchTarget || (d1 == searchTarget && ans > se) {
				ans = se
			}
		}

		if _, ok := in[searchTarget]; ok {
			return target
		}

		l, r := 0, len(next)-1
		get := -1
		for l <= r {
			mid := (r-l)/2 + l
			if searchTarget > next[mid] {
				get = mid
				l = mid + 1
				continue
			}
			r = mid - 1
		}
		if get == -1 {
			// select 0
			if ans == -1 {
				ans = se + next[0]
			} else {
				d1 := next[0] - searchTarget
				d2 := target - ans
				if d2 < 0 {
					d2 = -d2
				}
				if d2 > d1 || (d1 == d2 && ans > se+next[0]) {
					ans = se + next[0]
				}
			}
		} else {
			d1 := searchTarget - next[get]
			if ans == -1 {
				ans = se + next[get]
			} else {
				d2 := target - ans
				if d2 < 0 {
					d2 = -d2
				}
				if d2 > d1 || (d2 == d1 && ans > se+next[get]) {
					ans = se + next[get]
				}
			}

			if get < len(next)-1 {
				d1 := next[get+1] - searchTarget
				d2 := target - ans
				if d2 < 0 {
					d2 = -d2
				}
				if d2 > d1 || (d1 == d2 && ans > se+next[get+1]) {
					ans = se + next[get+1]
				}
			}
		}
	}
	return ans
}
