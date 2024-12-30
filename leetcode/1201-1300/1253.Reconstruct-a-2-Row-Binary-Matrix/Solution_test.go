package Solution

import (
	"reflect"
	"strconv"
	"testing"
)

func TestSolution(t *testing.T) {
	//	测试用例
	cases := []struct {
		name         string
		upper, lower int
		colsum       []int
		expect       [][]int
	}{
		{"TestCase1", 2, 1, []int{1, 1, 1}, [][]int{{1, 1, 0}, {0, 0, 1}}},
		{"TestCase2", 2, 3, []int{2, 2, 1, 1}, nil},
		{"TestCase3", 5, 5, []int{2, 1, 2, 0, 1, 0, 1, 2, 0, 1}, [][]int{{1, 1, 1, 0, 1, 0, 0, 1, 0, 0}, {1, 0, 1, 0, 0, 0, 1, 1, 0, 1}}},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			got := Solution(c.upper, c.lower, c.colsum)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v %v %v",
					c.expect, got, c.upper, c.lower, c.colsum)
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
