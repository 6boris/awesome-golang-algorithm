package Solution

import (
	"reflect"
	"strconv"
	"testing"
)

func TestSolution(t *testing.T) {
	//	测试用例
	cases := []struct {
		name    string
		price   []int
		special [][]int
		needs   []int
		expect  int
	}{
		{"TestCase1", []int{2, 5}, [][]int{
			{3, 0, 5}, {1, 2, 10},
		}, []int{3, 2}, 14},
		{"TestCase2", []int{2, 3, 4}, [][]int{
			{1, 1, 0, 4}, {2, 2, 1, 9},
		}, []int{1, 2, 1}, 11},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			got := Solution(c.price, c.special, c.needs)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v %v %v",
					c.expect, got, c.price, c.special, c.needs)
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
