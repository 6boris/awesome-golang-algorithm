package Solution

import "sort"

func Solution(s string, k int) int {
	if k == 0 {
		return 0
	}
	l := len(s)
	left, right := make([][3]int, l), make([][3]int, l)
	left[0][s[0]-'a'] = 1
	for i := 1; i < l; i++ {
		left[i] = left[i-1]
		left[i][s[i]-'a']++
	}
	right[l-1][s[l-1]-'a'] = 1
	for i := l - 2; i >= 0; i-- {
		right[i] = right[i+1]
		right[i][s[i]-'a']++
	}
	ans := -1
	for i := l - 1; i > 0; i-- {
		a, b, c := right[i][0], right[i][1], right[i][2]
		if a >= k && b >= k && c >= k {
			if ans == -1 || ans > l-i {
				ans = l - i
			}
			continue
		}
		idx := sort.Search(i, func(ii int) bool {
			return left[ii][0]+a >= k && left[ii][1]+b >= k && left[ii][2]+c >= k
		})
		if idx == i {
			continue
		}
		diff := l - i + idx + 1
		if ans == -1 || ans > diff {
			ans = diff
		}
	}

	for i := 0; i < l-1; i++ {
		a, b, c := left[i][0], left[i][1], left[i][2]
		if a >= k && b >= k && c >= k {
			if ans == -1 || ans > i+1 {
				ans = i + 1
			}
			continue
		}
		ll := l - i - 1
		start := i + 1
		idx := sort.Search(ll, func(ii int) bool {
			return right[start][0]+a >= k && right[start][1]+b >= k && right[start][2]+c >= k
		})
		if idx == ll {
			continue
		}
		diff := i + 1 + l - i
		if ans == -1 || ans > diff {
			ans = diff
		}
	}
	return ans
}
