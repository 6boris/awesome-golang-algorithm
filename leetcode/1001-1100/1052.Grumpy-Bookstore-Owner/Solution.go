package Solution

func Solution(customers []int, grumpy []int, minutes int) int {
	a := make([]int, len(customers))

	pre := 0
	total := 0

	for i := 0; i < len(customers); i++ {
		if grumpy[i] == 1 {
			pre += customers[i]
		} else {
			total += customers[i]
		}
		a[i] = pre
	}

	// 1,0,1,2,1,1,7,5
	// 0,1,0,1,0,1,0,1
	ans := 0
	for end := minutes - 1; end < len(customers); end++ {
		start := end + 1 - minutes
		get := a[end]
		if start > 0 {
			get -= a[start-1]
		}

		ans = max(ans, total+get)

	}

	return ans
}
