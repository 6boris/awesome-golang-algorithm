package Solution

func Solution(trips [][]int, capacity int) bool {
	maxDistance := 0
	for _, trip := range trips {
		if trip[2] > maxDistance {
			maxDistance = trip[2]
		}
	}

	distances := make([]int, maxDistance+1)
	for _, trip := range trips {
		for s := trip[1]; s < trip[2]; s++ {
			distances[s] += trip[0]
			if distances[s] > capacity {
				return false
			}
		}
	}
	return true
}
