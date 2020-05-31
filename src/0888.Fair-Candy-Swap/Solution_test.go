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
		A, B   []int
		expect []int
	}{
		{"TestCase", []int{1, 1}, []int{2, 2}, []int{1, 2}},
		{"TestCase", []int{1, 2}, []int{2, 3}, []int{1, 2}},
		{"TestCase", []int{2}, []int{1, 3}, []int{2, 3}},
		{"TestCase", []int{1, 2, 5}, []int{2, 4}, []int{5, 4}},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			got := Solution(c.A, c.B)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v %v",
					c.expect, got, c.A, c.B)
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
