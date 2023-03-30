package Solution

import (
	"strconv"
	"testing"
)

func TestSolution(t *testing.T) {
	//	测试用例
	cases := []struct {
		name   string
		count  int
		head   *ListNode
		expect []int
	}{
		{"TestCase1", 6, &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3}}}, []int{1, 1, 2, 2, 3, 3}},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			Solution(c.count, c.head)
		})
	}
}

// 压力测试
func BenchmarkSolution(b *testing.B) {
}

// 使用案列
func ExampleSolution() {
}
