package Solution

func TwoSum_1(nums []int, target int) []int {
	n := len(nums)

	for i, v := range nums {
		for j := i + 1; j < n; j++ {
			if v+nums[j] == target {
				return []int{i, j}
			}
		}
	}

	return nil
}

func TwoSum_2(nums []int, target int) []int {
	m := make(map[int]int, len(nums))

	for i, v := range nums {
		sub := target - v
		if j, ok := m[sub]; ok {
			return []int{j, i}
		} else {
			m[v] = i
		}
	}

	return nil
}

func TwoSum_3(nums []int, target int) []int {
	if nums == nil || len(nums) < 2 {
		return []int{-1, -1}
	}

	res := []int{-1, -1}

	//	MAP的KEY表示值，MAP的VAL表示nums的下标
	intMap := map[int]int{}
	for i := 0; i < len(nums); i++ {
		//	判断MAP中是否纯在一个Key满足 KEY + nums[i] = target
		//	如果满足则返回相关地址,不满足则将数组中的值PUSH到MAP
		if _, ok := intMap[target-nums[i]]; ok {
			res[0] = intMap[target-nums[i]]
			res[1] = i
			break
		}
		intMap[nums[i]] = i
	}
	return res
}
