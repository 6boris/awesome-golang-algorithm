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
		matrix [][]int
		k      int
		expect [][]int
	}{
		{"TestCase1", [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}, 1, [][]int{{12, 21, 16}, {27, 45, 33}, {24, 39, 28}}},
		{"TestCase2", [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}, 2, [][]int{{45, 45, 45}, {45, 45, 45}, {45, 45, 45}}},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			got := Solution(c.matrix, c.k)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with input-mat: %v, input-k: %v",
					c.expect, got, c.matrix, c.k)
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
