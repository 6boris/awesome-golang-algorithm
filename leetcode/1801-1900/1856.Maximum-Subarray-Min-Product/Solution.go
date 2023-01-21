package Solution

const mod1856 = 1000000007

func Solution(nums []int) int {
	l := len(nums)
	sum := make([]int64, l)
	sum[0] = int64(nums[0])
	for i := 1; i < l; i++ {
		sum[i] = int64(nums[i]) + sum[i-1]
	}

	nextSmaller := make([]int, l)
	prevSmaleer := make([]int, l)
	stack := make([]int, l)

	for i := 0; i < l; i++ {
		nextSmaller[i] = l
		prevSmaleer[i] = -1
	}

	stackIdx := -1
	for i := 0; i < l; i++ {
		for stackIdx >= 0 && nums[stack[stackIdx]] > nums[i] {
			nextSmaller[stack[stackIdx]] = i
			stackIdx--
		}
		stackIdx++
		stack[stackIdx] = i
	}

	stackIdx = -1
	for i := l - 1; i >= 0; i-- {
		for stackIdx >= 0 && nums[stack[stackIdx]] >= nums[i] {
			prevSmaleer[stack[stackIdx]] = i
			stackIdx--
		}
		stackIdx++
		stack[stackIdx] = i
	}

	n := int64(0)
	for i := 0; i < l; i++ {
		right := nextSmaller[i] - 1 // lastidx
		left := prevSmaleer[i]      //
		s := sum[right]
		if left != -1 {
			s -= sum[left]
		}

		tmp := int64(nums[i]) * s

		if n == -1 || n < tmp {
			n = tmp
		}
	}

	return int(n % mod1856)
}
