package Solution

type NumArray struct {
	Sum []int
	m   map[[2]int]int
}

//	暴力循环
func Constructor(nums []int) NumArray {
	var s NumArray
	s.Sum = nums
	return s
}

func (this *NumArray) SumRange(i int, j int) int {
	ans := 0
	for k := i; k <= j; k++ {
		ans += this.Sum[k]
	}
	return ans
}

//	利用 map 做缓存
func Constructor2(nums []int) NumArray {
	s := NumArray{m: map[[2]int]int{}}

	for i := 0; i < len(nums); i++ {
		sum := 0
		for j := i; j < len(nums); j++ {
			sum += nums[j]
			s.m[[2]int{i, j}] = sum
		}
	}
	return s
}

func (this *NumArray) SumRange2(i int, j int) int {
	return this.m[[2]int{i, j}]
}

//	DP
//	sum[i] = nums[0] +  nums[1] + nums[2]
//	sum[i+1] = sum[i] + nums[i+1]
func Constructor3(nums []int) NumArray {
	var res NumArray
	res.Sum = make([]int, len(nums)+1)
	for k, v := range nums {
		res.Sum[k+1] = res.Sum[k] + v
	}
	return res
}

func (this *NumArray) SumRange3(i int, j int) int {
	return this.Sum[j+1] - this.Sum[i]
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(i,j);
 */
