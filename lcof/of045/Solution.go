package Solution

import (
	"fmt"
	"sort"
	"strconv"
)

func minNumber(nums []int) string {
	ans := ""
	sort.Slice(nums, func(i, j int) bool {
		s1, s2 := fmt.Sprintf("%d%d", nums[i], nums[j]), fmt.Sprintf("%d%d", nums[j], nums[i])
		if s1 > s2 {
			return false
		}
		return true
	})
	for _, v := range nums {
		ans += strconv.Itoa(v)
	}
	return ans
}
