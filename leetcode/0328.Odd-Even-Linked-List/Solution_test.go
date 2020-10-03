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
		{"TestCase", MakeListNode([]int{1, 2, 3, 4, 5}), MakeListNode([]int{1, 3, 5, 2, 4})},
		{"TestCase", MakeListNode([]int{2, 1, 3, 5, 6, 4, 7}), MakeListNode([]int{2, 3, 6, 7, 1, 5, 4})},
		{"TestCase", MakeListNode([]int{}), MakeListNode([]int{})},
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

//	压力测试
func BenchmarkSolution(b *testing.B) {

}

//	使用案列
func ExampleSolution() {

}
