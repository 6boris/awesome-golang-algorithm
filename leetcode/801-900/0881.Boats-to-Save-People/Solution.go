package Solution

import "sort"

func Solution(people []int, limit int) int {
	sort.Ints(people)
	// b + d <= z, 可以得出a+d<=z,  且不可能出现b+c>z的情况
	ans := 0
	left, right := 0, len(people)-1
	for left <= right {
		ans++
		if people[left]+people[right] <= limit {
			left++
		}
		right--
	}

	return ans
}
