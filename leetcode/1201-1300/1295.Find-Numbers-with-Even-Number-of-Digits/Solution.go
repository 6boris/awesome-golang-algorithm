package Solution

import "strconv"

func Solution(nums []int) int {
	count := 0
	for i := 0; i < len(nums); i++ {
		digits := len(strconv.Itoa(nums[i]))
		if digits&1 == 0 {
			count++
		}
	}
	return count
}

func Solution1(nums []int) int {
	var isEven func(int) bool
	isEven = func(n int) bool {
		c := 0
		for n > 0 {
			n /= 10
			c++
		}
		return c&1 == 0
	}
	ans := 0
	for _, n := range nums {
		if isEven(n) {
			ans++
		}
	}
	return ans
}
