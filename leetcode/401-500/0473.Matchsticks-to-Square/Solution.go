package Solution

func Solution(matchsticks []int) bool {
	length := 0
	maxSide := 0
	for _, l := range matchsticks {
		if l > maxSide {
			maxSide = l
		}
		length += l
	}
	sideLength := length / 4
	if length%4 != 0 || len(matchsticks) < 4 || maxSide > sideLength {
		return false
	}
	used := make(map[int]struct{})
	return dfs473(4, 0, sideLength, 0, matchsticks, used)
}

// NOTE: dsf start, loop is always 4, square
func dfs473(loop, sum, target, idx int, nums []int, used map[int]struct{}) bool {
	if sum == target {
		return dfs473(loop-1, 0, target, 0, nums, used)
	}
	if loop == 1 {
		return true
	}
	for i := idx; i < len(nums); i++ {
		if _, ok := used[i]; !ok && sum+nums[i] <= target {
			used[i] = struct{}{}
			if dfs473(loop, sum+nums[i], target, i+1, nums, used) {
				return true
			}
			delete(used, i)
		}
	}
	return false
}
