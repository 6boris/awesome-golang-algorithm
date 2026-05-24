package Solution

func Solution(nums []int, limit int) int {
	n := len(nums)
	diff := make([]int, 2*limit+2)

	for i := 0; i < n/2; i++ {
		a, b := nums[i], nums[n-1-i]

		diff[2] += 2
		diff[2*limit+1] -= 2

		minVal := min(a, b)
		maxVal := max(a, b)
		diff[minVal+1] -= 1
		diff[maxVal+limit+1] += 1

		diff[a+b] -= 1
		diff[a+b+1] += 1
	}

	ans := n
	currentMoves := 0

	for s := 2; s <= 2*limit; s++ {
		currentMoves += diff[s]
		if currentMoves < ans {
			ans = currentMoves
		}
	}

	return ans
}
