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
		side   int
		points [][]int
		k      int
		expect int
	}{
		{"TestCase1", 2, [][]int{{0, 2}, {2, 0}, {2, 2}, {0, 0}}, 4, 2},
		{"TestCase2", 2, [][]int{{0, 0}, {1, 2}, {2, 0}, {2, 2}, {2, 1}}, 4, 1},
		{"TestCase3", 2, [][]int{{0, 0}, {0, 1}, {0, 2}, {1, 2}, {2, 0}, {2, 2}, {2, 1}}, 5, 1},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			got := Solution(c.side, c.points, c.k)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v %v %v",
					c.expect, got, c.side, c.points, c.k)
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
