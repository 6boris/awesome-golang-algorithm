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
		inputs *ListNode
		expect *ListNode
	}{
		{"TestCase1", &ListNode{Val: 1, Next: &ListNode{Val: 8, Next: &ListNode{Val: 9}}}, &ListNode{Val: 3, Next: &ListNode{Val: 7, Next: &ListNode{Val: 8}}}},
		{"TestCase2", &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9}}}, &ListNode{Val: 1, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 8}}}}},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			got := Solution(c.inputs)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v",
					c.expect, got, c.inputs)
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
