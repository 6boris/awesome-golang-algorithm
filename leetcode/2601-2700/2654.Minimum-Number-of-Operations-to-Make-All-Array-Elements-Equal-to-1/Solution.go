package Solution

func gcd2654(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func Solution(nums []int) int {
	l := len(nums)
	ones := 0
	// 有1就是直接响四周扩散就可以
	for _, n := range nums {
		if n == 1 {
			ones++
		}
	}
	if ones > 0 {
		return l - ones
	}

	// 没有，就需要看否构建出1
	minLenArrayWithOne := l + 1
	// 2, 2, 3, 4, 5
	for i := 0; i < l; i++ {
		gcdNum := nums[i]
		for j := i + 1; j < l; j++ {
			gcdNum = gcd2654(gcdNum, nums[j])
			if gcdNum == 1 {
				minLenArrayWithOne = min(minLenArrayWithOne, j-i+1)
				break
			}
		}
	}
	if minLenArrayWithOne == l+1 {
		return -1
	}
	// 构建minLenArrayWithOne里面的1，需要-1次，然后扩散到整个数组
	return minLenArrayWithOne - 1 + l - 1
}
