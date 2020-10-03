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
		A      []int
		K      int
		expect []int
	}{
		{"TestCase", []int{1, 2, 0, 0}, 34, []int{1, 2, 3, 4}},
		{"TestCase", []int{0}, 0, []int{0}},
		{"TestCase", []int{9, 9, 9, 9, 9}, 1, []int{1, 0, 0, 0, 0, 0}},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			got := Solution(c.A, c.K)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v %v",
					c.expect, got, c.A, c.K)
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
