package Solution

import (
	"reflect"
	"strconv"
	"testing"
)

func TestSolution(t *testing.T) {
	//	测试用例
	cases := []struct {
		name                    string
		robots, distance, walls []int
		expect                  int
	}{
		{"TestCase1", []int{4}, []int{3}, []int{1, 10}, 1},
		{"TestCase2", []int{10, 2}, []int{5, 1}, []int{5, 2, 7}, 3},
		{"TestCase3", []int{1, 2}, []int{100, 1}, []int{10}, 0},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			got := Solution(c.robots, c.distance, c.walls)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v %v %v",
					c.expect, got, c.robots, c.distance, c.walls)
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
