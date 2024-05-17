package Solution

import "sort"

func Solution(skill []int) int64 {
	sort.Ints(skill)
	cur := -1
	ans := int64(0)
	for s, e := 0, len(skill)-1; s < e; s, e = s+1, e-1 {
		if cur == -1 {
			cur = skill[s] + skill[e]
			ans += int64(skill[s]) * int64(skill[e])
			continue
		}
		if skill[s]+skill[e] != cur {
			return -1
		}
		ans += int64(skill[s]) * int64(skill[e])
	}
	return ans
}
