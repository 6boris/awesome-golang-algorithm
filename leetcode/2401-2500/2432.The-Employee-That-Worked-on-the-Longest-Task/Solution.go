package Solution

func Solution(n int, logs [][]int) int {
	slave, slaveCostTime := -1, -1
	costTime := 0
	for task := 0; task < len(logs); task++ {
		u, c := logs[task][0], logs[task][1]
		if x := c - costTime; x > slaveCostTime || (x == slaveCostTime && u < slave) {
			slave = u
			slaveCostTime = x
		}
		costTime = c
	}
	return slave
}
