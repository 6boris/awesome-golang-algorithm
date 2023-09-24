package Solution

import "strings"

func Solution(input string) int {

	ans := 0
	if !strings.Contains(input, "\n\t") {
		if !strings.Contains(input, ".") {
			return ans
		}

		size := 0
		for i := 0; i < len(input); i++ {
			if input[i] == '\n' {
				if size > ans {
					ans = size
				}
				size = 0
				continue
			}
			size++
		}
		if size > ans {
			return size
		}
		return ans
	}

	var dfs func(string, string, *int, int) int
	dfs = func(input, path string, index *int, tabs int) int {
		if *index >= len(input) {
			return tabs
		}

		l := len(input)
		for {
			// "a\n\tb.txt\na2\n\tb2.txt"
			i := *index
			// 读取一层dir
			tab := 0
			for ; i < len(input) && input[i] == '\t'; i++ {
				tab++
			}
			if tab != tabs+1 {
				return tab
			}

			size := 0
			start := i
			isEnd := false
			for ; i < l && input[i] != '\n'; i++ {
				size++
				if input[i] == '.' {
					isEnd = true
				}
			}
			addPath := path + "/" + input[start:i]
			*index = i + 1

			if isEnd {
				if r := len(addPath); r > ans {
					ans = r
				}
				continue
			}
			needTabs := dfs(input, addPath, index, tab)
			if needTabs != tabs+1 {
				return needTabs
			}
		}
	}

	start, end := 0, 0
	dirs := make([]string, 0)
	for ; end < len(input); end++ {
		if input[end] == '\n' {
			if end < len(input)-1 && input[end+1] != '\t' {
				dirs = append(dirs, input[start:end])
				start = end + 1
			}
		}
	}
	dirs = append(dirs, input[start:end])
	for _, baseInput := range dirs {
		index := 0
		for ; index < len(baseInput) && baseInput[index] != '\n'; index++ {
		}
		path := baseInput[:index]
		index++
		dfs(baseInput, path, &index, 0)
	}
	return ans
}
