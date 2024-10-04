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

func Solution1(skill []int) int64 {
	l := len(skill) / 2
	sum := 0
	for _, n := range skill {
		sum += n
	}
	if sum%l != 0 {
		return -1
	}
	target := sum / l
	count := make(map[int]int)
	for _, n := range skill {
		count[n]++
	}

	ans := int64(0)
	for k, c := range count {
		need := target - k
		if need == k && c&1 == 1 {
			return -1
		}
		if v, ok := count[need]; !ok || v != c {
			return -1
		}
		ans += int64(k) * int64(need) * int64(c)
	}
	return ans / 2
}
