package Solution

import (
	"reflect"
	"strconv"
	"testing"
)

func TestSolution(t *testing.T) {
	//	测试用例
	cases := []struct {
		name       string
		books      [][]int
		shelfWidth int
		expect     int
	}{
		{"TestCase1", [][]int{
			{1, 1},
			{2, 3},
			{2, 3},
		}, 4, 4},
		{"TestCase2", [][]int{
			{1, 1},
			{2, 3},
		}, 4, 3},
		{"TestCase3", [][]int{
			{1, 1},
			{2, 3},
			{2, 3},
			{1, 1},
			{1, 1},
			{1, 1},
			{1, 2},
		}, 4, 6},
		{"TestCase4", [][]int{
			{1, 2},
		}, 1, 2},
		{"TestCase5", [][]int{
			{1, 1},
			{2, 3},
			{2, 3},
			{1, 1},
			{1, 1},
			{1, 1},
			{1, 2},
			{3, 2},
		}, 5, 6},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			got := Solution(c.books, c.shelfWidth)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v %v",
					c.expect, got, c.books, c.shelfWidth)
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
