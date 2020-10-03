package Solution

func Solution(numbers []int, target int) []int {
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
	return []int{0, 0}
}
