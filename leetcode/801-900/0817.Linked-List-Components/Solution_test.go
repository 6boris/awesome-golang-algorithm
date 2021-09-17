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
		head   *ListNode
		G      []int
		expect int
	}{
		{"TestCase", MakeListNode([]int{0, 1, 2, 3}), []int{0, 1, 3}, 2},
		{"TestCase", MakeListNode([]int{0, 1, 2, 3, 4}), []int{0, 3, 1, 4}, 2},
	}

	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			got := Solution(c.head, c.G)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v %v",
					c.expect, got, c.head, c.G)
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
