package Solution

type deg struct {
	i, j, count int
}

func Solution(nums []int) int {
	hash, max := make(map[int]deg), -1
	for i := 0; i < len(nums); i++ {
		num := nums[i]
		if _, ok := hash[num]; ok {
			val := hash[num]
			val.count++
			if val.count > max {
				max = val.count
			}
			val.j = i
			hash[num] = val
		} else {
			val := deg{i, i, 1}
			if val.count > max {
				max = val.count
			}
			hash[num] = val
		}
	}
	min := 50000
	for _, val := range hash {
		v := val.j - val.i + 1
		if val.count == max && v > 0 && v < min {
			min = v
		}
	}
	return min
}
