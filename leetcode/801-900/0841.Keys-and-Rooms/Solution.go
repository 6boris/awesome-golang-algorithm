package Solution

func Solution(rooms [][]int) bool {
	visitedRomms := make([]bool, len(rooms))
	visitedRomms[0] = true
	canVisitAllRooms841(rooms[0], rooms, visitedRomms)
	for i := 0; i < len(rooms); i++ {
		if !visitedRomms[i] {
			return false
		}
	}
	return true
}

func canVisitAllRooms841(keys []int, rooms [][]int, visited []bool) {
	for _, k := range keys {
		if visited[k] {
			continue
		}
		visited[k] = true
		canVisitAllRooms841(rooms[k], rooms, visited)
	}
}
