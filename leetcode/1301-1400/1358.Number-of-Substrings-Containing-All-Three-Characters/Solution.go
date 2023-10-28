package Solution

import "sort"

func Solution(s string) int {
	need := [3][]int{}
	for i := 0; i < 3; i++ {
		need[i] = make([]int, 0)
	}
	ans := 0
	for i := len(s) - 1; i >= 0; i-- {
		need[s[i]-'a'] = append(need[s[i]-'a'], i)
	}
	for end := len(s) - 1; end > 1; end-- {
		if s[end] == 'a' {
			i1 := sort.Search(len(need[1]), func(i int) bool {
				return need[1][i] < end
			})
			i2 := sort.Search(len(need[2]), func(i int) bool {
				return need[2][i] < end
			})
			if i1 != len(need[1]) && i2 != len(need[2]) {
				x := need[1][i1]
				if need[2][i2] < x {
					x = need[2][i2]
				}
				ans += x + 1
			}
		}
		if s[end] == 'b' {
			i1 := sort.Search(len(need[0]), func(i int) bool {
				return need[0][i] < end
			})
			i2 := sort.Search(len(need[2]), func(i int) bool {
				return need[2][i] < end
			})
			if i1 != len(need[0]) && i2 != len(need[2]) {
				x := need[0][i1]
				if need[2][i2] < x {
					x = need[2][i2]
				}
				ans += x + 1
			}

		}
		if s[end] == 'c' {
			i1 := sort.Search(len(need[0]), func(i int) bool {
				return need[0][i] < end
			})
			i2 := sort.Search(len(need[1]), func(i int) bool {
				return need[1][i] < end
			})
			if i1 != len(need[0]) && i2 != len(need[1]) {
				x := need[0][i1]
				if need[1][i2] < x {
					x = need[1][i2]
				}
				ans += x + 1
			}
		}
	}
	return ans
}

func Solution1(s string) int {
	ans := 0

	chars := [3]int{}
	start := 0
	x := len(s)
	for i := 0; i < x; i++ {
		chars[s[i]-'a']++
		for chars[0] > 0 && chars[1] > 0 && chars[2] > 0 {
			ans += x - i
			chars[s[start]-'a']--
			start++
		}
	}
	return ans
}
