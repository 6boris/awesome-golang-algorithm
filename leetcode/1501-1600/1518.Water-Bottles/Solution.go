package Solution

func Solution(numBottles, numExchange int) int {
	emptyBottles := numBottles
	ans := numBottles
	for emptyBottles >= numExchange {
		bottles := emptyBottles / numExchange
		ans += bottles
		emptyBottles = emptyBottles%numExchange + bottles
	}
	return ans
}
