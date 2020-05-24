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
		input1 *ListNode
		input2 *ListNode
		expect *ListNode
	}{
		{"TestCase", MakeListNode([]int{7, 2, 4, 3}), MakeListNode([]int{5, 6, 4}), MakeListNode([]int{7, 8, 0, 7})},
		{"TestCase", MakeListNode([]int{}), MakeListNode([]int{1, 2, 3}), MakeListNode([]int{1, 2, 3})},
		{"TestCase", MakeListNode([]int{0, 0, 8}), MakeListNode([]int{9, 9, 2}), MakeListNode([]int{1, 0, 0, 0})},
		{"TestCase", MakeListNode([]int{}), MakeListNode([]int{}), MakeListNode([]int{})},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i+1), func(t *testing.T) {
			got := Solution(c.input1, c.input2)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v %v",
					c.expect, got, c.input1, c.input2)
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
