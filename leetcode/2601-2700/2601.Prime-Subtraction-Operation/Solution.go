package Solution

import "math"

func checkPrime(x int) bool {
	for i := 2; i <= int(math.Sqrt(float64(x))); i++ {
		if x%i == 0 {
			return false
		}
	}
	return true
}

func Solution(nums []int) bool {
	for i := 0; i < len(nums); i++ {
		var bound int
		if i == 0 {
			bound = nums[0]
		} else {
			bound = nums[i] - nums[i-1]
		}

		if bound <= 0 {
			return false
		}

		largestPrime := 0
		for j := bound - 1; j >= 2; j-- {
			if checkPrime(j) {
				largestPrime = j
				break
			}
		}

		nums[i] = nums[i] - largestPrime
	}
	return true
}
