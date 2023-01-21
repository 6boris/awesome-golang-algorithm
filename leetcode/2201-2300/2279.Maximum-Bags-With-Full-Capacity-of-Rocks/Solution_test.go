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
		a, b   []int
		c      int
		expect int
	}{
		{"TestCase1", []int{2, 3, 4, 5}, []int{1, 2, 4, 4}, 2, 3},
		{"TestCase2", []int{10, 2, 2}, []int{2, 2, 0}, 100, 3},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			got := Solution(c.a, c.b, c.c)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v %v %v",
					c.expect, got, c.a, c.b, c.c)
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
