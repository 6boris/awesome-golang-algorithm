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
		inputs int
		expect []int
	}{
		{"TestCase1", 1, []int{1}},
		{"TestCase2", 14, []int{1, 3, 4, 14}},
		{"TestCase3", 26, []int{1, 2, 6, 10, 26}},
		{"TestCase4", 1234, []int{1, 3, 4, 14, 19, 57, 77, 229, 308, 918, 1234}},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			got := Solution(c.inputs)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v",
					c.expect, got, c.inputs)
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
