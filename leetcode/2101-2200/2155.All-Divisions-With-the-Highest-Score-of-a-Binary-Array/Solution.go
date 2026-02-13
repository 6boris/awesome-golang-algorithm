package Solution

func Solution(nums []int) []int {
	var ret []int
	l := len(nums)
	ones := make([]int, l)
	ones[0] = nums[0]
	for i := 1; i < l; i++ {
		ones[i] = nums[i] + ones[i-1]
	}

	maxScore := 0
	var left, right, curScore int
	for i := 1; i < l; i++ {
		left = i - ones[i-1]
		right = ones[l-1] - ones[i-1]
		curScore = left + right
		if curScore > maxScore {
			maxScore = curScore
			ret = []int{i}
			continue
		}
		if curScore == maxScore {
			ret = append(ret, i)
		}
	}
	if ones[l-1] >= maxScore {
		if ones[l-1] == maxScore {
			ret = append(ret, 0)
		} else {
			ret = []int{0}
		}
		maxScore = ones[l-1]
	}
	left = l - ones[l-1]
	if left > maxScore {
		ret = []int{l}
	}
	if left == maxScore {
		ret = append(ret, l)
	}

	return ret
}
