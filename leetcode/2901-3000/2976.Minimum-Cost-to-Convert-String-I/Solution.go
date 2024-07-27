package Solution

const unreachable = 1 << 32

func Solution(source string, target string, original []byte, changed []byte, cost []int) int64 {
	cache := make([][]int64, 26)
	for i := 0; i < 26; i++ {
		cache[i] = make([]int64, 26)
		for j := 0; j < 26; j++ {
			cache[i][j] = unreachable
		}
		cache[i][i] = 0
	}
	// 定义图
	for i := 0; i < len(original); i++ {
		x, y := original[i]-'a', changed[i]-'a'
		cache[x][y] = min(cache[x][y], int64(cost[i]))
	}

	for k := 0; k < 26; k++ {
		for i := 0; i < 26; i++ {
			if cache[i][k] == unreachable {
				continue
			}
			for j := 0; j < 26; j++ {
				if cache[k][j] == unreachable {
					continue
				}
				cache[i][j] = min(cache[i][j], cache[i][k]+cache[k][j])
			}
		}
	}
	ans := int64(0)
	for i := 0; i < len(source); i++ {
		if source[i] != target[i] {
			r := cache[source[i]-'a'][target[i]-'a']
			if r == unreachable {
				return -1
			}
			ans += r
		}
	}
	return ans
}
