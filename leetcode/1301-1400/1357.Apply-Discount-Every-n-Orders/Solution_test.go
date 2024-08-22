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
		n, discount      int
		products, prices []int
		inputs           []input
		expect           []float64
	}{
		{"TestCase1", 3, 50, []int{1, 2, 3, 4, 5, 6, 7}, []int{100, 200, 300, 400, 300, 200, 100}, []input{
			{[]int{1, 2}, []int{1, 2}},
			{[]int{3, 7}, []int{10, 10}},
			{[]int{1, 2, 3, 4, 5, 6, 7}, []int{1, 1, 1, 1, 1, 1, 1}},
			{[]int{4}, []int{10}},
			{[]int{7, 3}, []int{10, 10}},
			{[]int{7, 5, 3, 1, 6, 4, 2}, []int{10, 10, 10, 9, 9, 9, 7}},
			{[]int{2, 3, 5}, []int{5, 3, 2}},
		}, []float64{500.0, 4000.0, 800.0, 4000.0, 4000.0, 7350.0, 2500.0}},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			got := Solution(c.n, c.discount, c.products, c.prices, c.inputs)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v",
					c.expect, got, c.inputs)
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
