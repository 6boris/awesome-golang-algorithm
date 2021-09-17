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
		expect []int
	}{
		{"TestCase", MakeListNode([]int{2, 1, 5}), []int{5, 5, 0}},
		{"TestCase", MakeListNode([]int{2, 7, 4, 3, 5}), []int{7, 0, 5, 5, 0}},
		{"TestCase", MakeListNode([]int{1, 7, 5, 1, 9, 2, 5, 1}), []int{7, 9, 9, 9, 0, 5, 0, 0}},
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
