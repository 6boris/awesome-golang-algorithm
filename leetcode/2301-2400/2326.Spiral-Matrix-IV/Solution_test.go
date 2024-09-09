package Solution

import (
	"reflect"
	"strconv"
	"testing"
)

func TestSolution(t *testing.T) {
	//	测试用例
	cases := []struct {
		name   string
		m, n   int
		head   *ListNode
		expect [][]int
	}{
		{"TestCase1", 3, 5, &ListNode{3, &ListNode{0, &ListNode{2, &ListNode{6, &ListNode{8, &ListNode{1, &ListNode{7, &ListNode{9, &ListNode{4, &ListNode{2, &ListNode{5, &ListNode{5, &ListNode{0, nil}}}}}}}}}}}}}, [][]int{
			{3, 0, 2, 6, 8},
			{5, 0, -1, -1, 1},
			{5, 2, 4, 9, 7},
		}},
		{"TestCase2", 1, 4, &ListNode{0, &ListNode{1, &ListNode{2, nil}}}, [][]int{
			{0, 1, 2, -1},
		}},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			got := Solution(c.m, c.n, c.head)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v %v %v",
					c.expect, got, c.m, c.n, c.head)
			}
		})
	}
}

// 压力测试
func BenchmarkSolution(b *testing.B) {
}

// 使用案列
func ExampleSolution() {
}
