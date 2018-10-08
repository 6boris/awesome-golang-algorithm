package Solution

//	HasMap解法
func twoSum(nums []int, target int) []int {
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

func twoSum1(nums []int, target int) []int {
	if nums == nil || len(nums) < 2 {
		return []int{-1, -1}
	}

	for k, v := range nums {
		for i := k + 1; i < len(nums); i++ {
			if v+nums[i] == target {
				return []int{k, i}
			}
		}
	}
	return nil
}
