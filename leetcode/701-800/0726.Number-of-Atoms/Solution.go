package Solution

import (
	"fmt"
	"sort"
	"strings"
)

type tmp726 struct {
	k string
	c int
}

func Solution(formula string) string {
	stack := make([]map[string]int, 0)
	count := make(map[string]int)
	i, l := 0, len(formula)
	for i < l {
		cur := formula[i]
		if cur >= 'A' && cur <= 'Z' {
			start := i
			i++
			// 判断后面是否还跟着小写字母
			for ; i < l && formula[i] >= 'a' && formula[i] <= 'z'; i++ {
			}
			// 一个元素
			key := formula[start:i]
			// 判断数字
			c := 0
			for ; i < l && formula[i] >= '0' && formula[i] <= '9'; i++ {
				c = c*10 + int(formula[i]-'0')
			}

			if c == 0 {
				c = 1
			}
			count[key] += c
			continue
		}
		if cur == '(' {
			stack = append(stack, count)
			count = map[string]int{}
			i++
			continue
		}
		pc := 0
		i++
		for ; i < l && formula[i] >= '0' && formula[i] <= '9'; i++ {
			pc = pc*10 + int(formula[i]-'0')
		}
		if pc == 0 {
			pc = 1
		}
		for k := range count {
			count[k] *= pc
		}
		if tl := len(stack); tl > 0 {
			top := stack[tl-1]
			stack = stack[:tl-1]
			for k, c := range top {
				count[k] += c
			}
		}

	}

	list := make([]tmp726, 0)
	for k, v := range count {
		list = append(list, tmp726{k, v})
	}
	sort.Slice(list, func(i, j int) bool {
		return list[i].k < list[j].k
	})
	buf := strings.Builder{}
	for _, i := range list {
		w := i.k
		if i.c != 1 {
			w += fmt.Sprintf("%d", i.c)
		}
		buf.WriteString(w)
	}
	return buf.String()
}
