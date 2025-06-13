package Solution

import (
	"reflect"
	"strconv"
	"testing"
)

func TestSolution(t *testing.T) {
	//	测试用例
	cases := []struct {
		name        string
		n, k        int
		invocations [][]int
		expect      []int
	}{
		{"TestCase1", 4, 1, [][]int{{1, 2}, {0, 1}, {3, 2}}, []int{0, 1, 2, 3}},
		{"TestCase2", 5, 0, [][]int{{1, 2}, {0, 2}, {0, 1}, {3, 4}}, []int{3, 4}},
		{"TestCase3", 3, 2, [][]int{{1, 2}, {0, 1}, {2, 0}}, nil},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			got := Solution(c.n, c.k, c.invocations)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v %v %v",
					c.expect, got, c.n, c.k, c.invocations)
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
