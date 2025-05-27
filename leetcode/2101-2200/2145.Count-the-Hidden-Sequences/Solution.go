package Solution

import "math"

func Solution(differences []int, lower int, upper int) int {
	sum := int64(0)
	var maxSum, minSum int64
	maxSum, minSum = int64(0), math.MaxInt64
	for _, n := range differences {
		sum += int64(n)
		maxSum = max(sum, maxSum)
		minSum = min(sum, minSum)
	}

	ans := 0
	il, ip := int64(lower), int64(upper)
	for i := lower; i <= upper; i++ {
		a := int64(i) + maxSum
		b := int64(i) + minSum
		if a >= il && a <= ip && b >= il && b <= ip {
			ans++
		}
	}
	return ans
}
