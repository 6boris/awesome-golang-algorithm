package Solution

func Solution(nums []int) int {
	n := len(nums)
	stack := make([]int, n)

	index := -1
	for i := 0; i < n; i++ {
		if index == -1 || nums[stack[index]] > nums[i] {
			index++
			stack[index] = i
		}
	}

	ans := 0

	for j := n - 1; j >= 0; j-- {
		for index != -1 && nums[stack[index]] <= nums[j] {
			ans = max(ans, j-stack[index])
			index--
		}
	}

	return ans
}
