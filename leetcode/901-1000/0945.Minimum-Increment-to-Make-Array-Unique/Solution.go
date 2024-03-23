package Solution

func Solution(nums []int) int {
	count := make(map[int]int)
	mn := -1
	for _, n := range nums {
		count[n]++
		if mn == -1 || mn > n {
			mn = n
		}
	}
	need := len(nums)
	steps := 0
	cur := mn
	for need > 0 {
		if count[cur] == 1 || count[cur] == 0 {
			if count[cur] != 0 {
				need--
			}
			cur++
			continue
		}
		next := cur + 1
		step := count[cur] - 1
		steps += step
		count[next] += step
		cur = next
		need--
	}
	return steps
}
