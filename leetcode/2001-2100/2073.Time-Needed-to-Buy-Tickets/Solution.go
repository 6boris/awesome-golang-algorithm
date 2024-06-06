package Solution

func Solution(tickets []int, k int) int {
	// 每个人1s的轮下来
	// 对于左侧的人，和右侧的人有差异，差异点在于：最后一轮到k后，右侧就不需要再走了
	// 所以最后从整体看就是，左侧:min(tickets[k], ticket[i]), 右侧就是min(tickets[k]-1,tickets[i])
	time := 0
	for i := 0; i < len(tickets); i++ {
		if i <= k {
			time += min(tickets[i], tickets[k])
		} else {
			time += min(tickets[k]-1, tickets[i])
		}
	}
	return time
}
