package Solution

var canUse = []byte{'a', 'b', 'c'}

func Solution(n int, k int) string {
	bs := make([]byte, n)
	var dfs func([]byte, byte, int)
	cnt := 0
	dfs = func(cur []byte, pre byte, index int) {
		if cnt > k {
			return
		}
		if index == n {
			cnt++
			if cnt == k {
				copy(bs, cur)
			}
			return
		}
		for _, b := range canUse {
			if b != pre {
				cur[index] = b
				dfs(cur, b, index+1)
			}
		}
	}
	arr := make([]byte, n)
	dfs(arr, byte(' '), 0)
	if cnt < k {
		return ""
	}
	return string(bs)
}
