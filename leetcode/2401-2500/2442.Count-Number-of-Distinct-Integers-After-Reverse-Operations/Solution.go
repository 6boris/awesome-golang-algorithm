package Solution

func rev(n int) int {
	x := 0
	for n > 0 {
		digit := n % 10
		n /= 10
		x = x*10 + digit
	}
	return x
}
func Solution(nums []int) int {
	set := make(map[int]struct{})
	for _, n := range nums {
		set[n] = struct{}{}
		set[rev(n)] = struct{}{}
	}
	return len(set)
}
