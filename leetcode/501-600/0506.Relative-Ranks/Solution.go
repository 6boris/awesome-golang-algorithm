package Solution

import (
	"sort"
	"strconv"
)

func Solution(score []int) []string {
	top := map[int]string{
		1: "Gold Medal",
		2: "Silver Medal",
		3: "Bronze Medal",
	}
	tmp := make([][2]int, len(score))
	for i := 0; i < len(score); i++ {
		tmp[i] = [2]int{i, score[i]}
	}
	sort.Slice(tmp, func(i, j int) bool {
		return tmp[i][1] > tmp[j][1]
	})
	ans := make([]string, len(score))
	for i := 0; i < len(tmp); i++ {
		if v, ok := top[i+1]; ok {
			ans[tmp[i][0]] = v
			continue
		}
		ans[tmp[i][0]] = strconv.Itoa(i + 1)
	}
	return ans
}
