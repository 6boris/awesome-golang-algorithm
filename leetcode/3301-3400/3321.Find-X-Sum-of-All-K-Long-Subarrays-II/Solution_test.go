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
		inputs []int
		k, x   int
		expect []int64
	}{
		{"TestCase1", []int{1, 1, 2, 2, 3, 4, 2, 3}, 6, 2, []int64{6, 10, 12}},
		{"TestCase2", []int{3, 8, 7, 8, 7, 5}, 2, 2, []int64{11, 15, 15, 15, 12}},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			got := Solution(c.inputs, c.k, c.x)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v %v %v",
					c.expect, got, c.inputs, c.k, c.x)
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
