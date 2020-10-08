package Solution

import (
	"testing"
)

func TestSolution(t *testing.T) {
	//	测试用例
	count := 9
	n1 := totalNQueens(count)
	n2 := totalNQueens3(count)
	n3 := totalNQueens2(count)
	t.Logf("n1=%d,n2=%d, n3=%d", n1, n2, n3)
}
