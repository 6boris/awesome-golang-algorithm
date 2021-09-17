package Solution

import (
	"reflect"
	"testing"
)

func TestSolution(t *testing.T) {
	//	测试用例
	cases := []struct {
		name   string
		head   *ListNode
		x      int
		expect *ListNode
	}{
		{"TestCase 1", MakeListNode([]int{1, 4, 3, 2, 5, 2}), 3, MakeListNode([]int{1, 2, 2, 4, 3, 5})},
		{"TestCase 2", MakeListNode([]int{}), 0, MakeListNode([]int{})},
		{"TestCase 3", MakeListNode([]int{1, 4, 3, 2, 5, 2}), 0, MakeListNode([]int{1, 4, 3, 2, 5, 2})},
	}

	//	开始测试
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := Solution(c.head, c.x)
			if !reflect.DeepEqual(ret, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v %v",
					c.expect, ret, c.head, c.x)
			}
		})
	}
}

//	压力测试
func BenchmarkSolution(b *testing.B) {
}

//	使用案列
func ExampleSolution() {
}
