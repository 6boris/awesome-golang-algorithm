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
		c, g   []int
		m      int
		expect int
	}{
		{"TestCase1", []int{1, 0, 1, 2, 1, 1, 7, 5}, []int{0, 1, 0, 1, 0, 1, 0, 1}, 3, 16},
		{"TestCase2", []int{1}, []int{0}, 1, 1},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			got := Solution(c.c, c.g, c.m)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v %v %v",
					c.expect, got, c.c, c.g, c.m)
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
