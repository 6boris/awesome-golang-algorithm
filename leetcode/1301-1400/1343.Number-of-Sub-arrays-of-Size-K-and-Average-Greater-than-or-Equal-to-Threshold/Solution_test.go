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
		inputs       []int
		k, threshold int
		expect       int
	}{
		{"TestCase1", []int{2, 2, 2, 2, 5, 5, 5, 8}, 3, 4, 3},
		{"TestCase2", []int{11, 13, 17, 23, 29, 31, 7, 5, 2, 3}, 3, 5, 6},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			got := Solution(c.inputs, c.k, c.threshold)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v %v %v",
					c.expect, got, c.inputs, c.k, c.threshold)
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
