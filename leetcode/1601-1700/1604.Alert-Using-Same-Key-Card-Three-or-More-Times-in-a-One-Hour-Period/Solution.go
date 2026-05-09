package Solution

import "sort"

func parseTime1604(now string) (int, int) {
	hour, minute := 0, 0
	i := 0
	for ; i < len(now) && now[i] != ':'; i++ {
		hour = hour*10 + int(now[i]-'0')
	}
	i++
	for ; i < len(now); i++ {
		minute = minute*10 + int(now[i]-'0')
	}
	return hour, minute
}

func Solution(keyName []string, keyTime []string) []string {
	var ret []string
	size := len(keyName)
	indies := make([]int, size)
	for i := range size {
		indies[i] = i
	}

	keyTimeInt := make([][2]int, size)
	for i := range keyTime {
		h, m := parseTime1604(keyTime[i])
		keyTimeInt[i] = [2]int{h, m}
	}

	sort.Slice(indies, func(i, j int) bool {
		ii, jj := indies[i], indies[j]
		if keyName[ii] == keyName[jj] {
			if keyTimeInt[ii][0] == keyTimeInt[jj][0] {
				return keyTimeInt[ii][1] < keyTimeInt[jj][1]
			}
			return keyTimeInt[ii][0] < keyTimeInt[jj][0]
		}
		return keyName[ii] < keyName[jj]
	})
	prev := " "
	added := make(map[string]struct{})

	start := 0
	for i := 0; i < size; i++ {
		if keyName[indies[i]] != prev {
			start = i
			prev = keyName[indies[i]]
			continue
		}
		// 看与start是否超过1小时
		diffH := keyTimeInt[indies[i]][0] - keyTimeInt[indies[start]][0]
		diffM := keyTimeInt[indies[i]][1] - keyTimeInt[indies[start]][1]
		minutes := diffM + diffH*60
		for minutes > 60 {
			start++
			diffH = keyTimeInt[indies[i]][0] - keyTimeInt[indies[start]][0]
			diffM = keyTimeInt[indies[i]][1] - keyTimeInt[indies[start]][1]
			minutes = diffM + diffH*60
		}
		if i-start+1 >= 3 {
			if _, ok := added[keyName[indies[i]]]; !ok {
				ret = append(ret, keyName[indies[i]])
				added[keyName[indies[i]]] = struct{}{}
			}
		}
	}
	return ret
}
