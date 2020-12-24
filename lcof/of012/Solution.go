package Solution

func exist(board [][]byte, word string) bool {
	m, n := len(board), len(board[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			//如果在数组中找得到第一个数，就执行下一步，否则返回false
			if dfs(board, i, j, word) {
				return true
			}
		}
	}
	return false
}
func dfs(board [][]byte, i, j int, word string) bool {
	//如果找到最后一个数，则返回true,搜索成功
	if len(word) == 0 {
		return true
	}
	//i,j的约束条件
	if i < 0 || j < 0 || i == len(board) || j == len(board[0]) {
		return false
	}

	//进入DFS深度优先搜索
	//先把正在遍历的该值重新赋值，如果在该值的周围都搜索不到目标字符，则再把该值还原
	//如果在数组中找到第一个字符，则进入下一个字符的查找
	if board[i][j] == word[0] {
		temp := board[i][j]
		board[i][j] = ' '
		//下面这个if语句，如果成功进入，说明找到该字符，然后进行下一个字符的搜索,直到所有的搜索都成功，
		//即k == len(word) - 1 的大小时，会返回true，进入该条件语句，然后返回函数true值。
		if dfs(board, i, j+1, word[1:]) ||
			dfs(board, i, j-1, word[1:]) ||
			dfs(board, i+1, j, word[1:]) ||
			dfs(board, i-1, j, word[1:]) {
			return true
		} else {
			//没有找到目标字符，还原之前重新赋值的board[i][j]
			board[i][j] = temp
		}
	}

	return false
}
