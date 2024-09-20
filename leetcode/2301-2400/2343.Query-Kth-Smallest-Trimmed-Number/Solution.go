package Solution

func radixSort2343(nums []string, digit int) [][]string {
	cache := make([][]string, digit)
	cur := len(nums[0]) - 1
	for i := 0; i < digit; i++ {
		buckets := [10][]string{}
		for _, num := range nums {
			key := num[cur] - '0'
			buckets[key] = append(buckets[key], num)
		}
		cur--
		index := 0
		for _, bucket := range buckets {
			for _, n := range bucket {
				nums[index] = n
				index++
			}
		}
		tmp := make([]string, len(nums))
		copy(tmp, nums)
		cache[i] = tmp
	}
	return cache
}

func Solution(nums []string, queries [][]int) []int {
	maxDigit := 0
	indies := make(map[string][]int)
	for i, n := range nums {
		if _, ok := indies[n]; !ok {
			indies[n] = make([]int, 0)
		}
		indies[n] = append(indies[n], i)
	}
	for _, q := range queries {
		maxDigit = max(maxDigit, q[1])
	}
	cache := radixSort2343(nums, maxDigit)
	ans := make([]int, len(queries))
	for i, q := range queries {
		sortedArray := cache[q[1]-1]
		index := 0
		target := sortedArray[q[0]-1]
		for pre := q[0] - 2; pre >= 0 && sortedArray[pre] == target; pre-- {
			index++
		}
		ans[i] = indies[target][index]
	}

	return ans
}
