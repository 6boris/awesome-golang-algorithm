package Solution

import (
	"reflect"
	"strconv"
	"testing"
)

func TestSolution(t *testing.T) {
	//	测试用例
	cases := []struct {
		name     string
		days     int
		meetings [][]int
		expect   int
	}{
		{"TestCase1", 10, [][]int{{5, 7}, {1, 3}, {9, 10}}, 2},
		{"TestCase2", 5, [][]int{{2, 4}, {1, 3}}, 1},
		{"TestCase3", 6, [][]int{{1, 6}}, 0},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			got := Solution(c.days, c.meetings)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v %v",
					c.expect, got, c.days, c.meetings)
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
