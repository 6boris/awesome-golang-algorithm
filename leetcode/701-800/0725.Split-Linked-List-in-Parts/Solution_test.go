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
		input2 int
		expect []*ListNode
	}{
		{"TestCase", MakeListNode([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}), 3,
			[]*ListNode{MakeListNode([]int{1, 2, 3, 4}), MakeListNode([]int{5, 6, 7}), MakeListNode([]int{8, 9, 10})}},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
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
