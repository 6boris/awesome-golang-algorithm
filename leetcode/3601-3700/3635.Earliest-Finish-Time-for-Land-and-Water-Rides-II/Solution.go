package Solution

func Solution(landStartTime, landDuration, waterStartTime, waterDuration []int) int {
	ans := 0x7fffffff

	minLandEnd := 0x7fffffff
	for i := range landStartTime {
		minLandEnd = min(minLandEnd, landStartTime[i]+landDuration[i])
	}
	for i := range waterStartTime {
		ans = min(ans, max(minLandEnd, waterStartTime[i])+waterDuration[i])
	}

	minWaterEnd := 0x7fffffff
	for i := range waterStartTime {
		minWaterEnd = min(minWaterEnd, waterStartTime[i]+waterDuration[i])
	}
	for i := range landStartTime {
		ans = min(ans, max(minWaterEnd, landStartTime[i])+landDuration[i])
	}

	return ans
}
