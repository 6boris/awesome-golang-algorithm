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
		q, w   []int
		k      int
		expect float64
	}{
		{"TestCase1", []int{10, 20, 5}, []int{70, 50, 30}, 2, 105.0000},
		{"TestCase2", []int{3, 1, 10, 10, 1}, []int{4, 8, 2, 2, 7}, 3, 30.666666666666664},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			got := Solution(c.q, c.w, c.k)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v %v %v",
					c.expect, got, c.q, c.w, c.k)
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
