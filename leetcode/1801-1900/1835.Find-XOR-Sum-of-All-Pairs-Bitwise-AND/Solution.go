package Solution

func Solution(arr1 []int, arr2 []int) int {
	xor2 := 0
	for _, n := range arr2 {
		xor2 ^= n
	}
	ans := 0
	for _, n := range arr1 {
		ans ^= (n & xor2)
	}
	return ans
}
