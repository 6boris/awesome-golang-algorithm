package Solution

func Solution(nums []int) [][]int {
	ans := make([][]int, 0)
	l := len(nums)
	a := make([]int, l+1)
	loop := 0
	for _, n := range nums {
		a[n]++
		if a[n] > loop {
			loop = a[n]
		}
	}
	for index := 0; loop > 0; index, loop = index+1, loop-1 {
		ans = append(ans, []int{})
		for i := 1; i <= l; i++ {
			if a[i] > 0 {
				ans[index] = append(ans[index], i)
				a[i]--
			}
		}
	}
	return ans
}
