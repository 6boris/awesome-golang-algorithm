package Solution

func Solution(colors string, neededTime []int) int {
	sum := neededTime[0]
	ans := 0
	mm := neededTime[0]
	c := 1
	for idx := 1; idx < len(colors); idx++ {
		if colors[idx] == colors[idx-1] {
			sum += neededTime[idx]
			if mm < neededTime[idx] {
				mm = neededTime[idx]
			}
			c++
			continue
		}
		if c > 1 {
			ans += sum - mm
		}
		mm = neededTime[idx]
		sum = neededTime[idx]
		c = 1
	}
	if c > 1 {
		ans += sum - mm
	}
	return ans
}
