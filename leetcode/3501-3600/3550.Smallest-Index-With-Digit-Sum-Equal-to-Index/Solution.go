package Solution

func Solution(nums []int) int {
	sum := 0
	for i := range nums {
		sum = 0
		for n := nums[i]; n > 0; n /= 10 {
			sum += n % 10
		}
		if sum == i {
			return i
		}
	}
	return -1
}
