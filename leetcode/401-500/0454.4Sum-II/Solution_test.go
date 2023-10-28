package Solution

import (
	"reflect"
	"strconv"
	"testing"
)

func TestSolution(t *testing.T) {
	//	测试用例
	cases := []struct {
		name           string
		n1, n2, n3, n4 []int
		expect         int
	}{
		{"TestCase1", []int{1, 2}, []int{-2, -1}, []int{-1, 2}, []int{0, 2}, 2},
		{"TestCase2", []int{0}, []int{0}, []int{0}, []int{0}, 1},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			got := Solution(c.n1, c.n2, c.n3, c.n4)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v %v %v %v",
					c.expect, got, c.n1, c.n2, c.n3, c.n4)
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
