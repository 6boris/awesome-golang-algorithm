package Solution

import (
	"sort"
)

func Solution(hand []int, groupSize int) bool {
	l := len(hand)
	if l%groupSize != 0 {
		return false
	}
	if groupSize == 1 {
		return true
	}

	keys := make([]int, 0)
	keyCount := make(map[int]int)
	for _, n := range hand {
		keyCount[n]++
		if keyCount[n] == 1 {
			keys = append(keys, n)
		}
	}
	sort.Ints(keys)

	idx := 1
	count := keyCount[keys[0]]
	nextIdx := -1
	keyCount[keys[0]] = 0
	used := 1
	l -= count

	for idx < len(keys) {
		if keys[idx]-keys[idx-1] != 1 {
			return false
		}
		keyCount[keys[idx]] -= count
		l -= count
		used++
		if keyCount[keys[idx]] < 0 {
			return false
		}
		if keyCount[keys[idx]] > 0 && nextIdx == -1 {
			nextIdx = idx
		}

		if used != groupSize {
			idx++
			continue
		}
		if nextIdx == -1 {
			if idx == len(keys)-1 {
				break
			}
			idx++
			count = keyCount[keys[idx]]
			nextIdx = -1
			keyCount[keys[idx]] = 0
			l -= count
			used = 1
			idx++
			continue
		}

		idx = nextIdx + 1
		count = keyCount[keys[nextIdx]]
		keyCount[keys[nextIdx]] = 0
		nextIdx = -1
		l -= count
		used = 1
	}
	return l == 0 && used == groupSize
}
