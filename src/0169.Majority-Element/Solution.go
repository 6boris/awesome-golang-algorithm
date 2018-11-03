package Solution

//	解法1： 循环计数，代码比较简单，LeetCode上的Case都是有Majority的，所以这里确实很多判断
func majorityElement1(nums []int) int {
	major, count := nums[0], 0

	for _, v := range nums {
		if count == 0 {
			count++
			major = v
		} else if major == v {
			count++
		} else {
			count--
		}
	}

	return major
}

//	解法2: 2次循环
func majorityElement2(nums []int) int {
	major, count := nums[0], 0

	for _, v1 := range nums {
		tmpCount := 0
		for _, v2 := range nums {
			if v2 == v1 {
				tmpCount++
			}
		}
		if tmpCount > count {
			count = tmpCount
			major = v1
		}
	}

	return major
}

//	解法3: Map
func majorityElement3(nums []int) int {
	major, count := nums[0], 0
	majorMap := make(map[int]int)

	for _, v := range nums {
		majorMap[v]++
	}

	for i, v := range majorMap {
		if v > count {
			major = i
			count = v
		}
	}

	return major
}

//	解法4: 先排序再依次比对
func majorityElement4(nums []int) int {
	//	将原数组排序
	nums = InsertSort(nums)

	major, count := nums[0], 0

	for _, v := range nums {
		if count == 0 {
			count++
			major = v
		} else if major == v {
			count++
		} else {
			count--
		}
	}

	return major
}

// 插入排序(Insert Sort)
func InsertSort(arr []int) []int {
	var i, j, tmp int
	for i = 1; i < len(arr); i++ {
		tmp = arr[i]
		for j = i; j > 0 && arr[j-1] > tmp; j-- {
			arr[j] = arr[j-1]
		}
		arr[j] = tmp
	}
	return arr
}

//	解法5: 分治算法
func majorityElement5(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	n := len(nums)

	if n%2 == 1 {
		count := 1

		for i := 0; i < n-1; i++ {
			if nums[n-1] == nums[i] {
				count++
			}
		}
		if count > n/2 {
			return nums[n-1]
		}
	}
	numsTmp := []int{}

	for i := 0; i < n/2; i++ {
		if nums[2*i] == nums[2*i+1] {
			numsTmp = append(numsTmp, nums[2*i])
		}
	}

	return majorityElement1(numsTmp)
}
