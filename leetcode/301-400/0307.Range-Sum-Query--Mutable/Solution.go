package Solution

type NumArray struct {
	raw []int
}

func Constructor307(nums []int) NumArray {
	for idx := 1; idx < len(nums); idx++ {
		nums[idx] += nums[idx-1]
	}
	return NumArray{raw: append([]int{0}, nums...)}
}

func (this *NumArray) Update(index int, val int) {
	source := this.raw[index+1] - this.raw[index]
	diff := val - source
	for idx := index + 1; idx < len(this.raw); idx++ {
		this.raw[idx] += diff
	}
}

func (this *NumArray) SumRange(left int, right int) int {
	return this.raw[right+1] - this.raw[left]
}

// 这道题看解法，是线段树
func Solution(nums []int, operations []string, optNums [][]int) []int {
	o := Constructor307(nums)
	ans := make([]int, 0)
	for idx, opt := range operations {
		if opt == "update" {
			o.Update(optNums[idx][0], optNums[idx][1])
			continue
		}
		if opt == "sumRange" {
			exp := o.SumRange(optNums[idx][0], optNums[idx][1])
			ans = append(ans, exp)
		}
	}
	return ans
}
