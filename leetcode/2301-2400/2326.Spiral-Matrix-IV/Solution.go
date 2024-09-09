package Solution

type ListNode struct {
	Val  int
	Next *ListNode
}

func Solution(m int, n int, head *ListNode) [][]int {
	ans := make([][]int, m)
	for i := range m {
		ans[i] = make([]int, n)
		for j := range n {
			ans[i][j] = -1
		}
	}

	x, y := 0, 0
	width := 0
	for head != nil {
		for ; y < n-width && head != nil; y, head = y+1, head.Next {
			ans[x][y] = head.Val
		}
		x, y = x+1, y-1

		for ; x < m-width && head != nil; x, head = x+1, head.Next {
			ans[x][y] = head.Val
		}
		x, y = x-1, y-1

		for ; y >= width && head != nil; y, head = y-1, head.Next {
			ans[x][y] = head.Val
		}
		x, y = x-1, y+1
		for ; x > width && head != nil; x, head = x-1, head.Next {
			ans[x][y] = head.Val
		}
		width++
		x, y = width, width
	}
	return ans
}
