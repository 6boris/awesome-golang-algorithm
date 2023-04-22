package Solution

import (
	"reflect"
	"strconv"
	"testing"
)

func TestSolution(t *testing.T) {
	//	测试用例
	cases := []struct {
		name         string
		baseCosts    []int
		toppingCosts []int
		target       int
		expect       int
	}{
		{"TestCase1", []int{1, 7}, []int{3, 4}, 10, 10},
		{"TestCase2", []int{2, 3}, []int{3, 4, 100}, 18, 17},
		{"TestCase3", []int{3, 10}, []int{2, 5}, 9, 8},
		{"TestCase4", []int{9, 10, 1}, []int{1, 8, 8, 1, 1, 8}, 8, 7},
		{"TestCase5", []int{3}, []int{3}, 9, 9},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			got := Solution(c.baseCosts, c.toppingCosts, c.target)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v %v %v",
					c.expect, got, c.baseCosts, c.toppingCosts, c.target)
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
