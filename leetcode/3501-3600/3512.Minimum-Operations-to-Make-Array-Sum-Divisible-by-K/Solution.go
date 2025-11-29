package Solution

func Solution(nums []int, k int) int {
	sum := 0
	for _, n := range nums {
		sum += n
	}
	return sum % k
}
