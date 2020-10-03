package Solution

func Solution(nums []int) int {
	ans := len(nums)
	for i, num := range nums {
		ans ^= i ^ num
	}
	return ans
}
