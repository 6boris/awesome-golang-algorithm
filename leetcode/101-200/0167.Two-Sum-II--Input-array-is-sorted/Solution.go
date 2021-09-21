package Solution

func twoSum_1(numbers []int, target int) []int {
	for i := 0; i < len(numbers); i++ {
		low, high := i+1, len(numbers)-1
		for low <= high {
			mid := (high-low)/2 + low
			if numbers[mid] == target-numbers[i] {
				return []int{i + 1, mid + 1}
			} else if numbers[mid] > target-numbers[i] {
				high = mid - 1
			} else {
				low = mid + 1
			}
		}
	}
	return []int{-1, -1}
}

func twoSum_2(numbers []int, target int) []int {
	i, j := 0, len(numbers)-1
	for i < j {
		sum := numbers[i] + numbers[j]
		if sum < target {
			i++
		} else if sum > target {
			j--
		} else {
			return []int{i + 1, j + 1}
		}
	}
	return []int{-1, -1}
}
