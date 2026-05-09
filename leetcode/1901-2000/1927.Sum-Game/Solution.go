package Solution

func Solution(num string) bool {
	n := len(num)
	leftSum, rightSum := 0, 0
	leftQ, rightQ := 0, 0

	for i := 0; i < n/2; i++ {
		if num[i] == '?' {
			leftQ++
		} else {
			leftSum += int(num[i] - '0')
		}
	}

	for i := n / 2; i < n; i++ {
		if num[i] == '?' {
			rightQ++
		} else {
			rightSum += int(num[i] - '0')
		}
	}

	diffS := leftSum - rightSum
	diffQ := rightQ - leftQ

	return diffS*2 != diffQ*9
}
