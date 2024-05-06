package Solution

type NumArray struct {
	data []int
}

func Constructor(nums []int) NumArray {
	for i := 1; i < len(nums); i++ {
		nums[i] += nums[i-1]
	}
	return NumArray{data: nums}
}

func (this *NumArray) SumRange(left int, right int) int {
	del := 0
	if left > 0 {
		del = this.data[left-1]
	}
	return this.data[right] - del
}

func Solution(nums []int, operations [][2]int) []int {
	ans := make([]int, len(operations))
	c := Constructor(nums)
	for i, op := range operations {
		ans[i] = c.SumRange(op[0], op[1])
	}
	return ans
}
