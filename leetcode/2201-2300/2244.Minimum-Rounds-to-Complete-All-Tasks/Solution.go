package Solution

func Solution(tasks []int) int {
	taskLevel := make(map[int]int)
	for _, t := range tasks {
		taskLevel[t]++
	}
	rounds := 0

	for _, count := range taskLevel {
		if count == 1 {
			return -1
		}
		cond := count % 3
		if cond == 0 {
			rounds += count / 3
			continue
		}
		// 4(2+2), 7(3+2+2), 10(3+3+2+2), 13(3+3+3+2+2)
		if cond == 1 {
			rounds += (count-4)/3 + 2
			continue
		}
		// 5(3+2), 8(3+3+2), 11(3+3+3+2)
		rounds += count/3 + 1
	}
	return rounds
}
