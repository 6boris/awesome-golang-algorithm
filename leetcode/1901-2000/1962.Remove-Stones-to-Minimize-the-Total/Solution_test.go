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
		k      int
		expect int
	}{
		{"TestCase1", []int{5, 4, 9}, 2, 12},
		{"TestCase2", []int{4, 3, 6, 7}, 3, 12},
		{"TestCase3", []int{10, 88, 74, 63, 28, 99, 131, 74, 97, 201, 383, 274}, 6, 865},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			got := Solution(c.inputs, c.k)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v %v",
					c.expect, got, c.inputs, c.k)
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
