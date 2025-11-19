package Solution

import (
	"sort"
)

type ElementInfo struct {
	Value int
	Freq  int
}

func Solution(nums []int, k int, x int) []int {
	n := len(nums)
	if k == 0 || n < k {
		return []int{}
	}

	answer := make([]int, n-k+1)

	for i := 0; i <= n-k; i++ {
		freqMap := make(map[int]int)
		for j := i; j < i+k; j++ {
			freqMap[nums[j]]++
		}

		var elements []ElementInfo
		for val, freq := range freqMap {
			elements = append(elements, ElementInfo{
				Value: val,
				Freq:  freq,
			})
		}

		sort.Slice(elements, func(a, b int) bool {
			if elements[a].Freq != elements[b].Freq {
				return elements[a].Freq > elements[b].Freq
			}
			return elements[a].Value > elements[b].Value
		})

		currentXSum := 0

		limit := x
		if len(elements) < x {
			limit = len(elements)
		}

		for idx := 0; idx < limit; idx++ {
			element := elements[idx]
			currentXSum += element.Value * element.Freq
		}

		answer[i] = currentXSum
	}

	return answer
}
