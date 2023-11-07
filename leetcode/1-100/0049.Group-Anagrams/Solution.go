package Solution

import (
	"fmt"
	"sort"
)

func groupAnagrams(nums []string) [][]string {
	ans := [][]string{{}}
	m := map[string][]string{}
	for _, s := range nums {
		tmp := []byte(s)
		sort.Slice(tmp, func(i, j int) bool {
			return tmp[i] > tmp[j]
		})
		m[string(tmp)] = append(m[string(tmp)], s)
	}
	fmt.Println(m)
	for _, v := range m {
		ans = append(ans, v)
	}
	return ans
}
