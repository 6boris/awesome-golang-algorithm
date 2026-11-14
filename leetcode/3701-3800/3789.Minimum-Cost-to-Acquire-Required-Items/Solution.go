package Solution

func Solution(cost1 int, cost2 int, costBoth int, need1 int, need2 int) int64 {
	// type1:1, type2:1, type3:各自+1
	// cost1    cost2    costBoth
	i64cost1 := int64(cost1)
	i64cost2 := int64(cost2)
	i64both := int64(costBoth)
	i64need1 := int64(need1)
	i64need2 := int64(need2)

	// 1, 各自单买
	case1 := i64cost1*i64need1 + i64cost2*i64need2

	// costBoth * min(need1, need2) + extra
	extra2 := max(i64need2-i64need1, 0)
	case2 := i64need1*i64both + extra2*i64cost2

	extra1 := max(i64need1-i64need2, 0)
	case3 := i64need2*i64both + extra1*i64cost1

	// costBoth * max(need1, need2)
	return min(case1, case2, case3, i64both*max(i64need1, i64need2))
}
