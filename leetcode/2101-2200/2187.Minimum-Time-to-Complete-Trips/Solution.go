package Solution

import "sort"

func Solution(time []int, totalTrips int) int64 {
	ans := int64(0)
	sort.Ints(time)

	// giao
	minTime, maxTime := int64(1), int64(time[len(time)-1]*totalTrips)
	for minTime <= maxTime {
		midTime := (maxTime-minTime)/2 + minTime
		done := 0
		for i := 0; i < len(time); i++ {
			t := midTime / int64(time[i])
			done += int(t)
			if done >= totalTrips {
				break
			}
		}
		if done >= totalTrips {
			ans = midTime
			maxTime = midTime - 1
		} else {
			minTime = midTime + 1
		}

	}
	/*
		for ; ; ans++ {
			done := 0
			for i := 0; i < len(time); i++ {
				tmp := ans / int64(time[i])
				done += int(tmp)
				if done >= totalTrips {
					return ans
				}
			}
		}
	*/
	return ans
}
