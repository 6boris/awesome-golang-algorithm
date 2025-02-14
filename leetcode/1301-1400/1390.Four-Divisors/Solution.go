package Solution

import "math"

func check(n int) (bool, int) {
	r := 0
	ans := 0
	for i := 1; i <= int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			another := n / i
			r++
			ans += i
			if another != i {
				r++
				ans += another
			}
			if r > 4 {
				return false, 0
			}
		}
	}
	return r == 4, ans
}

func Solution(nums []int) int {
	ans := 0
	for _, n := range nums {
		ok, v := check(n)
		if !ok {
			continue
		}
		ans += v
	}
	return ans
}
