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
		n1, n2, n3 []int
		expect []int
	}{
		{"TestCase1", []int{1, 1, 3, 2}, []int{2, 3}, []int{3}, []int{2 ,3}},
		{"TestCase2", []int{3, 1}, []int{2, 3}, []int{1, 2}, []int{1, 2, 3}},
		{"TestCase3", []int{1, 2, 2}, []int{4, 3, 3}, []int{5}, []int{}},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			got := Solution(c.n1, c.n2, c.n3)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v %v %v",
					c.expect, got, c.n1, c.n2, c.n3)
			}
		})
	}
}

//	压力测试
func BenchmarkSolution(b *testing.B) {
}

//	使用案列
func ExampleSolution() {
}
