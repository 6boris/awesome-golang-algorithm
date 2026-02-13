package Solution

import (
	"strconv"
)

func Solution(names []string) []string {
	l := len(names)
	ret := make([]string, l)
	cache := make(map[string]int)
	var tmpName, suffix string
	for idx, name := range names {
		if _, ok := cache[name]; !ok {
			cache[name] = 1 // 下一个从1开始
			ret[idx] = name
			continue
		}
		nextNumber := cache[name]
		for {
			suffix = strconv.Itoa(nextNumber)
			tmpName = name + "(" + suffix + ")"
			//fmt.Printf("testing %s\n", tmpName)
			_, ok := cache[tmpName]
			if ok {
				nextNumber++
				continue
			}
			ret[idx] = tmpName
			//fmt.Printf("--tmpName: %s\n", tmpName)
			cache[tmpName] = 1
			nextNumber++
			cache[name] = nextNumber
			break
		}
	}
	return ret
}
