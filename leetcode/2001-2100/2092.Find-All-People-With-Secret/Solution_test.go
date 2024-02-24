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
		n           int
		meetings    [][]int
		firstPerson int
		expect      []int
	}{
		{"TestCase1", 6, [][]int{
			{1, 2, 5}, {2, 3, 8}, {1, 5, 10},
		}, 1, []int{0, 1, 2, 3, 5}},
		{"TestCase2", 4, [][]int{
			{3, 1, 3}, {1, 2, 2}, {0, 3, 3},
		}, 3, []int{0, 1, 3}},
		{"TestCase3", 5, [][]int{
			{3, 4, 2}, {1, 2, 1}, {2, 3, 1},
		}, 1, []int{0, 1, 2, 3, 4}},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			got := Solution(c.n, c.meetings, c.firstPerson)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v %v %v",
					c.expect, got, c.n, c.meetings, c.firstPerson)
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
