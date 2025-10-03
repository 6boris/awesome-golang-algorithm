package Solution

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type transaction struct {
	name         string
	time, amount int
	city         string
}

func (tr transaction) String() string {
	return fmt.Sprintf("%s,%d,%d,%s", tr.name, tr.time, tr.amount, tr.city)
}

func Solution(transactions []string) []string {
	list := make([]transaction, len(transactions))
	for i, trans := range transactions {
		items := strings.Split(trans, ",")
		time, _ := strconv.Atoi(items[1])
		amount, _ := strconv.Atoi(items[2])
		list[i] = transaction{
			name:   items[0],
			time:   time,
			amount: amount,
			city:   items[3],
		}
	}
	var ret []string
	sort.Slice(list, func(i, j int) bool {
		// 感觉时间排序
		a, b := list[i], list[j]
		if a.time == b.time {
			return a.name < b.name
		}
		return a.time < b.time
	})

	group := make(map[string][]transaction)
	for _, item := range list {
		if _, ok := group[item.name]; !ok {
			group[item.name] = make([]transaction, 0)
		}
		group[item.name] = append(group[item.name], item)
	}
	for _, trans := range group {
		invalid := make([]bool, len(trans))
		for i := 0; i < len(trans); i++ {

			for pre := i - 1; pre >= 0; pre-- {
				if trans[i].city == trans[pre].city {
					continue
				}
				if trans[i].time <= trans[pre].time+60 {
					invalid[i] = true
					invalid[pre] = true
				}
			}
			if trans[i].amount > 1000 {
				invalid[i] = true
				continue
			}
		}
		for i := range trans {
			if invalid[i] {
				ret = append(ret, trans[i].String())
			}
		}
	}
	return ret
}
