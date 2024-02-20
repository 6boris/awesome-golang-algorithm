package Solution

import (
	"strconv"
	"strings"
)

func Solution(n int, logs []string) []int {
	ans := make([]int, n)
	stack := [][2]int{}
	for _, log := range logs {
		part := strings.Split(log, ":")
		id, _ := strconv.Atoi(part[0])
		timestamp, _ := strconv.Atoi(part[2])
		if part[1] == "start" {
			l := len(stack)
			if l != 0 {
				top := stack[l-1]
				ans[top[0]] += timestamp - top[1]
			}
			stack = append(stack, [2]int{id, timestamp})
			continue
		}

		l := len(stack)
		top := stack[l-1]
		ans[top[0]] += timestamp - top[1] + 1
		stack = stack[:l-1]
		if l = len(stack); l != 0 {
			stack[l-1][1] = timestamp + 1
		}
	}
	return ans
}
