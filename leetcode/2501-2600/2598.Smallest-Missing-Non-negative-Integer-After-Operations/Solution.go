package Solution

func Solution(nums []int, value int) int {
	mp := make([]int, value)
	for _, x := range nums {
		v := ((x % value) + value) % value
		mp[v]++
	}
	mex := 0
	for mp[mex%value] > 0 {
		mp[mex%value]--
		mex++
	}
	return mex
}
