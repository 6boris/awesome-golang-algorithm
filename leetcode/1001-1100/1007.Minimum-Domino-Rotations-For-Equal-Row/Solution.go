package Solution

func topCanImpl(elem int, tops, bottoms []int) (bool, int) {
	count := 0
	for idx := 1; idx < len(tops); idx++ {
		if tops[idx] == elem {
			continue
		}
		if bottoms[idx] != elem {
			return false, count
		}
		count++
	}

	return true, count
}

func bottomCanImpl(elem int, tops, bottoms []int) (bool, int) {
	count := 0
	for idx := 1; idx < len(bottoms); idx++ {
		if bottoms[idx] == elem {
			continue
		}
		if tops[idx] != elem {
			return false, count
		}
		count++
	}

	return true, count
}

func Solution(tops []int, bottoms []int) int {
	minCount := -1
	if ttOk, ttCount := topCanImpl(tops[0], tops, bottoms); ttOk && (minCount == -1 || minCount > ttCount) {
		minCount = ttCount
	}
	if tbOk, tbCount := topCanImpl(bottoms[0], tops, bottoms); tbOk && (minCount == -1 || minCount > tbCount+1) {
		minCount = tbCount + 1
	}
	if bbOk, bbCount := bottomCanImpl(bottoms[0], tops, bottoms); bbOk && (minCount == -1 || minCount > bbCount) {
		minCount = bbCount
	}
	if btOk, btCount := bottomCanImpl(tops[0], tops, bottoms); btOk && (minCount == -1 || minCount > btCount+1) {
		minCount = btCount + 1
	}
	return minCount
}
