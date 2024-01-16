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
		n1, n2 []int
		expect int
	}{
		{"TestCase1", []int{1, 7, 5}, []int{2, 3, 5}, 3},
		{"TestCase2", []int{2, 4, 6, 8, 10}, []int{2, 4, 6, 8, 10}, 0},
		{"TestCase3", []int{1, 10, 4, 4, 2, 7}, []int{9, 3, 5, 1, 7, 4}, 20},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			got := Solution(c.n1, c.n2)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v %v",
					c.expect, got, c.n1, c.n2)
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
