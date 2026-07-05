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
		grid   [][]int
		health int
		expect bool
	}{
		{"TestCase1", [][]int{
			{0, 1, 0, 0, 0},
			{0, 1, 0, 1, 0},
			{0, 0, 0, 1, 0},
		}, 1, true},
		{"TestCase2", [][]int{
			{0, 1, 1, 0, 0, 0},
			{1, 0, 1, 0, 0, 0},
			{0, 1, 1, 1, 0, 1},
			{0, 0, 1, 0, 1, 0},
		}, 3, false},
		{"TestCase3", [][]int{
			{1, 1, 1}, {1, 0, 1}, {1, 1, 1},
		}, 5, true},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			got := Solution(c.grid, c.health)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v %v",
					c.expect, got, c.grid, c.health)
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
