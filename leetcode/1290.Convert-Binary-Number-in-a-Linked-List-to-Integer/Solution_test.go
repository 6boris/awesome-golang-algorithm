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
		expect int
	}{
		{"TestCase", &ListNode{1, &ListNode{0, &ListNode{1, nil}}}, 5},
		{"TestCase", &ListNode{0, nil}, 0},
		{"TestCase", &ListNode{1, &ListNode{1, &ListNode{1, nil}}}, 7},
		{"TestCase", &ListNode{0, &ListNode{1, nil}}, 1},
		{"TestCase", &ListNode{1, &ListNode{0, nil}}, 2},
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
