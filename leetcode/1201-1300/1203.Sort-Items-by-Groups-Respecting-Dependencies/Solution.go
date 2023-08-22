package Solution

import "sort"

func Solution(n int, m int, group []int, beforeItems [][]int) []int {
	groupItems := make(map[int][]int)
	item2Group := make(map[int]int)
	for idx, g := range group {
		if g == -1 {
			g = m + idx
		}
		item2Group[idx] = g
		if _, ok := groupItems[g]; !ok {
			groupItems[g] = make([]int, 0)
		}
		groupItems[g] = append(groupItems[g], idx)
	}

	/*
		for g, items := range groupItems {
			fmt.Printf("group: %d -- %v\n", g, items)
		}
		for item, g := range item2Group {
			fmt.Printf("itme: %d -- %d\n", item, g)
		}
	*/
	groupItemOrder := make(map[int]map[int][]int)
	groupItemIn := make(map[int]map[int]int)

	groupOrder := make(map[int][]int)
	groupIn := make(map[int]int)
	for idx, before := range beforeItems {
		idxGroup := item2Group[idx]

		for _, item := range before {
			itemGroup := item2Group[item]
			if itemGroup == idxGroup {
				if _, ok := groupItemOrder[idxGroup]; !ok {
					groupItemOrder[idxGroup] = make(map[int][]int)
				}
				if _, ok := groupItemOrder[idxGroup][item]; !ok {
					groupItemOrder[idxGroup][item] = make([]int, 0)
				}
				groupItemOrder[idxGroup][item] = append(groupItemOrder[idxGroup][item], idx)

				if _, ok := groupItemIn[idxGroup]; !ok {
					groupItemIn[idxGroup] = make(map[int]int)
				}
				groupItemIn[idxGroup][idx]++
				continue
			}
			if _, ok := groupOrder[itemGroup]; !ok {
				groupOrder[itemGroup] = make([]int, 0)
			}
			groupOrder[itemGroup] = append(groupOrder[itemGroup], idxGroup)
			groupIn[idxGroup]++
		}
	}

	/*
		fmt.Printf("groupItemOrder; %+v\n", groupItemOrder)
		fmt.Printf("groupOrder: %+v\n", groupOrder)
		fmt.Printf("groupIn: %+v\n", groupIn)
	*/
	var groupTopology func(int) bool
	groupTopology = func(groupID int) bool {
		items := groupItems[groupID]
		if len(items) == 1 {
			return true
		}
		newOrder := make([]int, 0)
		relation := groupItemOrder[groupID]
		in := groupItemIn[groupID]

		queue := make([]int, 0)
		used := 0
		for _, item := range items {
			if _, ok := in[item]; !ok {
				queue = append(queue, item)
				newOrder = append(newOrder, item)
				used++
			}
		}
		if len(queue) == 0 {
			return false
		}
		// The sorting is added here for test cases, and can be removed from the actual commit code
		sort.Ints(newOrder)
		for len(queue) > 0 {
			n := make([]int, 0)
			for _, i := range queue {
				for _, rel := range relation[i] {
					in[rel]--
					if in[rel] == 0 {
						n = append(n, rel)
					}
				}
			}
			if len(n) == 0 && used != len(items) {
				return false
			}
			queue = n
			used += len(n)
			// The sorting is added here for test cases, and can be removed from the actual commit code
			sort.Ints(n)
			newOrder = append(newOrder, n...)
		}

		groupItems[groupID] = newOrder
		return true
	}
	// 先检查每个组是否可以满足要求
	for groupID := range groupItems {
		if !groupTopology(groupID) {
			//fmt.Printf("group; %d, don't", groupID)
			return nil
		}
	}

	queue := make([]int, 0)
	newGroupOrder := make([]int, 0)
	used := 0
	for i := range groupItems {
		if _, ok := groupIn[i]; !ok {
			queue = append(queue, i)
			//fmt.Printf("add group: %d\n", i)
			newGroupOrder = append(newGroupOrder, i)
			used++
		}
	}
	if used == 0 {
		return nil
	}
	// The sorting is added here for test cases, and can be removed from the actual commit code
	sort.Ints(newGroupOrder)

	for len(queue) > 0 {
		n := make([]int, 0)
		for _, item := range queue {
			for _, rel := range groupOrder[item] {
				groupIn[rel]--
				if groupIn[rel] == 0 {
					n = append(n, rel)
				}
			}
		}
		if len(n) == 0 && used != len(groupItems) {
			return nil
		}
		queue = n
		// The sorting is added here for test cases, and can be removed from the actual commit code
		sort.Ints(n)
		newGroupOrder = append(newGroupOrder, n...)
		used += len(n)
	}

	ans := make([]int, 0)
	for _, g := range newGroupOrder {
		ans = append(ans, groupItems[g]...)
	}

	return ans
}
