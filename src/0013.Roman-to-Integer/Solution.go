package Solution

func romanToInt(s string) int {
	//	初始化一个罗马字符的Map
	romanMap := map[string]int{}
	romanMap["I"] = 1
	romanMap["V"] = 5
	romanMap["X"] = 10
	romanMap["L"] = 50
	romanMap["C"] = 100
	romanMap["D"] = 500
	romanMap["M"] = 1000

	MapIndex := string(s[len(s)-1])
	sum := romanMap[MapIndex]
	for i := len(s) - 2; i >= 0; i-- {
		//	左小右大
		if romanMap[string(s[i])] < romanMap[string(s[i+1])] {
			sum -= romanMap[string(s[i])]
			//	左大右小
		} else {
			sum += romanMap[string(s[i])]
		}
	}

	return sum
}
