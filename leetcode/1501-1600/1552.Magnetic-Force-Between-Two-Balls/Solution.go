package Solution

import "sort"

func canPlaceBalls(x int, position []int, m int) bool {
	prevBallPos := position[0]
	ballsPlaced := 1

	for i := 1; i < len(position) && ballsPlaced < m; i++ {
		currPos := position[i]
		if currPos-prevBallPos >= x {
			ballsPlaced++
			prevBallPos = currPos
		}
	}

	return ballsPlaced == m
}

func Solution(position []int, m int) int {
	answer := 0
	n := len(position)
	sort.Ints(position)

	low := 1
	high := int(position[n-1]) / (m - 1)
	for low <= high {
		mid := low + (high-low)/2
		if canPlaceBalls(mid, position, m) {
			answer = mid
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return answer
}
