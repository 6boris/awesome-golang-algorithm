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
		k, w   int
		p, c   []int
		expect int
	}{
		{"TestCase1", 2, 0, []int{1, 2, 3}, []int{0, 1, 1}, 4},
		{"TestCase2", 3, 0, []int{1, 2, 3}, []int{0, 1, 2}, 6},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			got := Solution(c.k, c.w, c.p, c.c)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v %v %v %v",
					c.expect, got, c.k, c.w, c.p, c.c)
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
