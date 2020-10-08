package Solution

var maps = []string{"abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return nil
	}
	var res []string
	var str string
	dfs(digits, &res, str, 0)
	return res
}

func dfs(digits string, res *[]string, str string, start int) {
	if start == len(digits) {
		*res = append(*res, str)
		return
	}
	mapStr := maps[digits[start]-'2']
	for i := 0; i < len(mapStr); i++ {
		dfs(digits, res, str+string(mapStr[i]), start+1)
	}
}

var numberMap = map[byte][]string{
	'2': []string{"a", "b", "c"},
	'3': []string{"d", "e", "f"},
	'4': []string{"g", "h", "i"},
	'5': []string{"j", "k", "l"},
	'6': []string{"m", "n", "o"},
	'7': []string{"p", "q", "r", "s"},
	'8': []string{"t", "u", "v"},
	'9': []string{"w", "x", "y", "z"},
}

func depthSolution(s string) []string {
	if len(s) == 1 {
		return numberMap[s[0]]
	}
	returnArr := []string{}
	for _, e := range numberMap[s[0]] {
		for _, e1 := range depthSolution(s[1:]) {
			returnArr = append(returnArr, e+e1)
		}
	}
	return returnArr
}
func breadthSolution(s string) []string {
	returnQueue := numberMap[(s[0])]
	position := 1
	for position < len(s) {
		newQueue := []string{}
		for _, e := range returnQueue {
			for _, e1 := range numberMap[s[position]] {
				newQueue = append(newQueue, e+e1)
			}
		}
		returnQueue = newQueue
		position++
	}
	return returnQueue
}
