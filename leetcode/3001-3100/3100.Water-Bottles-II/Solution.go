package Solution

func Solution(numBottles, numExchange int) int {
	ans := numBottles
	drink := 0
	empty := numBottles
	for {
		if empty >= numExchange {
			drink++
			empty -= numExchange
			numExchange++
			continue
		}
		empty += drink
		ans += drink
		drink = 0
		if empty < numExchange {
			break
		}
	}
	return ans + drink
}
