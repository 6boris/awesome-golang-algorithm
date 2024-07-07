package Solution

import (
	"sort"
	"strconv"
)

func Solution(orders [][]string) [][]string {
	foodHeaders := make([]string, 0)
	foodHeaderIndex := make(map[string]struct{})
	tables := make([]string, 0)
	tableFood := make(map[string]map[string]int)

	for _, order := range orders {
		t, f := order[1], order[2]
		if _, ok := foodHeaderIndex[f]; !ok {
			foodHeaders = append(foodHeaders, f)
			foodHeaderIndex[f] = struct{}{}
		}

		if _, ok := tableFood[t]; !ok {
			tables = append(tables, t)
			tableFood[t] = make(map[string]int)
		}
		tableFood[t][f]++
	}
	sort.Strings(foodHeaders)
	foodHeaders = append([]string{"Table"}, foodHeaders...)
	sort.Slice(tables, func(i, j int) bool {
		a, _ := strconv.Atoi(tables[i])
		b, _ := strconv.Atoi(tables[j])
		return a < b
	})

	ans := make([][]string, 0)
	ans = append(ans, foodHeaders)
	for _, t := range tables {
		line := make([]string, len(foodHeaders))
		line[0] = t
		for i := 1; i < len(line); i++ {
			line[i] = strconv.Itoa(tableFood[t][foodHeaders[i]])
		}
		ans = append(ans, line)
	}
	return ans
}
