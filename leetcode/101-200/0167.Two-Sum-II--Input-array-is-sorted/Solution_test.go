package Solution

import (
	"reflect"
	"strconv"
	"testing"
)

func TestSolution(t *testing.T) {
	//	测试用例
	cases := []struct {
		name    string
		numbers []int
		target  int
		expect  []int
	}{
		{"TestCase", []int{2, 7, 11, 15}, 9, []int{1, 2}},
		{"TestCase", []int{2, 7, 11, 15}, 22, []int{2, 4}},
		{"TestCase", []int{2, 7, 11, 15}, 26, []int{3, 4}},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			got := Solution(c.numbers, c.target)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v %v",
					c.expect, got, c.numbers, c.target)
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
