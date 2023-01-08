package Solution

const mod907 = 1000000007

func Solution(arr []int) int {
	l := len(arr)
	nextSmaller := make([]int, l)
	prevSmaller := make([]int, l)

	stack := make([]int, l)
	for i := 0; i < l; i++ {
		nextSmaller[i] = l
		prevSmaller[i] = -1
	}
	startIdx := -1
	for i := 0; i < l; i++ {
		for startIdx >= 0 && arr[stack[startIdx]] > arr[i] {
			nextSmaller[stack[startIdx]] = i
			startIdx--
		}
		startIdx++
		stack[startIdx] = i
	}
	startIdx = -1
	for i := l - 1; i >= 0; i-- {
		for startIdx >= 0 && arr[stack[startIdx]] >= arr[i] {
			prevSmaller[stack[startIdx]] = i
			startIdx--
		}
		startIdx++
		stack[startIdx] = i
	}

	ans := 0
	for i := 0; i < l; i++ {
		right := nextSmaller[i] - i
		left := i - prevSmaller[i] - 1
		tmp := right + left*right
		ans += tmp * arr[i]
	}
	return ans % mod907
}
