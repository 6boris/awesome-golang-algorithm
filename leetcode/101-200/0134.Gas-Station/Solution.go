package Solution

func Solution(gas []int, cost []int) int {
	l := len(gas)
	tank, total, target := 0, 0, 0
	for i := 0; i < l; i++ {
		tank += gas[i] - cost[i]
		if tank < 0 {
			target = i + 1
			total += tank
			tank = 0
		}
	}

	if tank+total < 0 {
		return -1
	}
	return target
}
