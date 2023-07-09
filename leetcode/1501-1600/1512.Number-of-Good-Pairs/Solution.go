package Solution

func Solution(nums []int) int {
	count := make([]int, 101)
	for _, item := range nums {
		count[item]++
	}
	ans := 0
	for i := 1; i <= 100; i++ {
		if count[i] > 1 {
			n := (count[i] - 1) * count[i]
			ans += n / 2
		}
	}
	return ans
}
