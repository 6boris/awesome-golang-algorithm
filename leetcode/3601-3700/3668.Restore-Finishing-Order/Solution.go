package Solution

func Solution(order []int, friends []int) []int {
	friendsMap := make(map[int]struct{})
	for _, n := range friends {
		friendsMap[n] = struct{}{}
	}
	index := 0
	for _, n := range order {
		if _, ok := friendsMap[n]; ok {
			friends[index] = n
			index++
		}
	}
	return friends
}
