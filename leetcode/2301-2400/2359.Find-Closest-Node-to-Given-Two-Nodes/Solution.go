package Solution

func Solution(edges []int, node1 int, node2 int) int {

	var canReach func(int, int, map[int]int)
	canReach = func(startNode, l int, path map[int]int) {
		if _, ok := path[startNode]; ok {
			return
		}
		path[startNode] = l
		next := edges[startNode]
		if next == -1 {
			return
		}
		canReach(next, l+1, path)
	}
	node1Path := make(map[int]int)
	node2Path := make(map[int]int)
	canReach(node1, 1, node1Path)
	canReach(node2, 1, node2Path)

	min, nodeIndex := -1, -1

	for n1, dis1 := range node1Path {
		if dis2, ok := node2Path[n1]; ok {
			max := dis1
			if dis2 > max {
				max = dis2
			}
			if min == -1 || min > max || (min == max && n1 < nodeIndex) {
				min = max
				nodeIndex = n1
			}
		}
	}
	return nodeIndex
}
