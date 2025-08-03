package Solution

import (
	"reflect"
	"strconv"
	"testing"
)

func TestSolution(t *testing.T) {
	//	测试用例
	cases := []struct {
		name             string
		basket1, basket2 []int
		expect           int64
	}{
		{"TestCase1", []int{4, 2, 2, 2}, []int{1, 4, 1, 2}, 1},
		{"TestCase2", []int{2, 3, 4, 1}, []int{1, 2, 5, 1}, -1},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			got := Solution(c.basket1, c.basket2)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v %v",
					c.expect, got, c.basket1, c.basket2)
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
