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
		logs   [][]int
		k      int
		expect []int
	}{
		{"TestCase1", [][]int{{0, 5}, {1, 2}, {0, 2}, {0, 5}, {1, 3}}, 5, []int{0, 2, 0, 0, 0}},
		{"TestCase2", [][]int{{1, 1}, {2, 2}, {2, 3}}, 4, []int{1, 1, 0, 0}},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			got := Solution(c.logs, c.k)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v %v",
					c.expect, got, c.logs, c.k)
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
