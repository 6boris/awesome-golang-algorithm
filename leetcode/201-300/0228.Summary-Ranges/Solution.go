package Solution

import (
	"fmt"
)

func Solution(nums []int) []string {
	result := make([]string, 0)
	if len(nums) == 0 {
		return result
	}
	start, end := 0, 1
	for ; end < len(nums); end++ {
		if nums[end]-nums[end-1] == 1 {
			continue
		}

		writeObj := fmt.Sprintf("%d", nums[start])
		if end-start != 1 {
			writeObj = fmt.Sprintf("%d->%d", nums[start], nums[end-1])
		}
		result = append(result, writeObj)
		start = end
	}
	writeObj := fmt.Sprintf("%d", nums[start])
	if end-start != 1 {
		writeObj = fmt.Sprintf("%d->%d", nums[start], nums[end-1])
	}
	result = append(result, writeObj)
	return result
}
