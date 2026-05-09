package Solution

import "sort"

func Solution(robots []int, distance []int, walls []int) int {
	n := len(robots)
	left := make([]int, n)
	right := make([]int, n)
	num := make([]int, n)
	robotsToDistance := make(map[int]int)

	for i := 0; i < n; i++ {
		robotsToDistance[robots[i]] = distance[i]
	}

	sort.Ints(robots)
	sort.Ints(walls)

	for i := 0; i < n; i++ {
		pos1 := sort.SearchInts(walls, robots[i]+1)

		var leftPos int
		if i >= 1 {
			leftBound := max(robots[i]-robotsToDistance[robots[i]], robots[i-1]+1)
			leftPos = sort.SearchInts(walls, leftBound)
		} else {
			leftPos = sort.SearchInts(walls, robots[i]-robotsToDistance[robots[i]])
		}
		left[i] = pos1 - leftPos

		var rightPos int
		if i < n-1 {
			rightBound := min(robots[i]+robotsToDistance[robots[i]], robots[i+1]-1)
			rightPos = sort.SearchInts(walls, rightBound+1)
		} else {
			rightPos = sort.SearchInts(walls, robots[i]+robotsToDistance[robots[i]]+1)
		}
		pos2 := sort.SearchInts(walls, robots[i])
		right[i] = rightPos - pos2

		if i == 0 {
			continue
		}
		pos3 := sort.SearchInts(walls, robots[i-1])
		num[i] = pos1 - pos3
	}

	subLeft, subRight := left[0], right[0]
	for i := 1; i < n; i++ {
		currentLeft := max(subLeft+left[i], subRight-right[i-1]+min(left[i]+right[i-1], num[i]))
		currentRight := max(subLeft+right[i], subRight+right[i])
		subLeft, subRight = currentLeft, currentRight
	}

	return max(subLeft, subRight)
}
