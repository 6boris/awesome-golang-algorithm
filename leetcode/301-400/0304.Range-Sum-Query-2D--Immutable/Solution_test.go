package Solution

import (
	"reflect"
	"strconv"
	"testing"
)

func TestSolution(t *testing.T) {
	//	测试用例
	cases := []struct {
		name        string
		matrix      [][]int
		coordinates [][4]int
		expect      []int
	}{
		{"TestCase1", [][]int{
			{3, 0, 1, 4, 2},
			{5, 6, 3, 2, 1},
			{1, 2, 0, 1, 5},
			{4, 1, 0, 1, 7},
			{1, 0, 3, 0, 5},
		}, [][4]int{
			{2, 1, 4, 3},
			{1, 1, 2, 2},
			{1, 2, 2, 4},
			{0, 0, 0, 0},
			{4, 4, 4, 4},
		}, []int{8, 11, 12, 3, 5}},
		{"TestCase2", [][]int{
			{-4, -5},
		}, [][4]int{
			{0, 0, 0, 0},
			{0, 0, 0, 1},
			{0, 1, 0, 1},
		}, []int{-4, -9, -5}},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			got := Solution(c.matrix, c.coordinates)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v %v",
					c.expect, got, c.matrix, c.coordinates)
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
