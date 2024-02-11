package Solution

func Solution(routes [][]int, source int, target int) int {
	if source == target {
		return 0
	}
	// 记录每个station在某个车的什么index上面
	stationInCar := make(map[int]map[int]int)
	// 记录哪些车的某个站会经过target
	targetInCar := make(map[int]struct{})
	sourceInCar := make([][3]int, 0)
	carCrossStation := make(map[int][]int)
	visited := make(map[int]map[int]struct{})

	for car, route := range routes {
		visited[car] = make(map[int]struct{})

		stationInCar[car] = make(map[int]int)
		for idx, station := range route {
			stationInCar[car][station] = idx
			if station == target {
				targetInCar[car] = struct{}{}
			}
			if station == source {
				sourceInCar = append(sourceInCar, [3]int{car, station, 1})
			}
			if _, ok := carCrossStation[station]; !ok {
				carCrossStation[station] = make([]int, 0)
			}
			carCrossStation[station] = append(carCrossStation[station], car)
		}
	}
	ans := -1
	for len(sourceInCar) > 0 {
		nq := make([][3]int, 0)
		for _, s := range sourceInCar {
			curCar, curStation, curCost := s[0], s[1], s[2]
			if _, ok := targetInCar[curCar]; ok {
				if ans == -1 || curCost < ans {
					ans = curCost
				}
				continue
			}
			for _, next := range carCrossStation[curStation] {
				if next == curCar {
					curIndex := stationInCar[curCar][curStation]
					nextIndex := (curIndex + 1) % len(stationInCar[curCar])
					if _, ok := visited[curCar][routes[curCar][nextIndex]]; ok {
						continue
					}
					visited[curCar][routes[curCar][nextIndex]] = struct{}{}
					nq = append(nq, [3]int{curCar, routes[curCar][nextIndex], curCost})
				} else {
					if _, ok := visited[next][curStation]; ok {
						continue
					}
					visited[next][curStation] = struct{}{}
					nq = append(nq, [3]int{next, curStation, curCost + 1})
				}
			}
		}
		sourceInCar = nq
	}
	return ans
}
