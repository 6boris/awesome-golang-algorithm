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
		k      int
		expect [][]int
	}{
		{"TestCase", [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}, 1, [][]int{{9, 1, 2}, {3, 4, 5}, {6, 7, 8}}},
		{"TestCase", [][]int{{3, 8, 1, 9}, {19, 7, 2, 5}, {4, 6, 11, 10}, {12, 0, 21, 13}}, 4, [][]int{{12, 0, 21, 13}, {3, 8, 1, 9}, {19, 7, 2, 5}, {4, 6, 11, 10}}},
		{"TestCase", [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}, 9, [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			got := Solution(c.grid, c.k)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v %v",
					c.expect, got, c.grid, c.k)
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
