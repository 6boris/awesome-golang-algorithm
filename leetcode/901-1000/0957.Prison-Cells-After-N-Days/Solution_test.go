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
		cells  []int
		n      int
		expect []int
	}{
		{"TestCase1", []int{0, 1, 0, 1, 1, 0, 0, 1}, 7, []int{0, 0, 1, 1, 0, 0, 0, 0}},
		{"TestCase2", []int{1, 0, 0, 1, 0, 0, 1, 0}, 1000000000, []int{0, 0, 1, 1, 1, 1, 1, 0}},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			got := Solution(c.cells, c.n)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v %v",
					c.expect, got, c.cells, c.n)
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
