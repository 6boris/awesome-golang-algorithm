package Solution

import (
	"reflect"
	"strconv"
	"testing"
)

func TestSolution(t *testing.T) {
	//	测试用例
	cases := []struct {
		name      string
		n         int
		relations [][]int
		time      []int
		expect    int
	}{
		{"TestCase1", 3, [][]int{
			{1, 3}, {2, 3},
		}, []int{3, 2, 5}, 8},
		{"TestCase2", 5, [][]int{
			{1, 5}, {2, 5}, {3, 5}, {3, 4}, {4, 5},
		}, []int{1, 2, 3, 4, 5}, 12},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			got := Solution(c.n, c.relations, c.time)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v %v %v",
					c.expect, got, c.n, c.relations, c.time)
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
