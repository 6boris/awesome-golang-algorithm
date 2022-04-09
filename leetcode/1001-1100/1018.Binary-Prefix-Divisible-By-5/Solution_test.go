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
		inputs []int
		expect []bool
	}{
		{"TestCase", []int{0}, []bool{true}},
		{"TestCase", []int{1}, []bool{false}},
		{"TestCase", []int{0, 1, 1}, []bool{true, false, false}},
		{"TestCase", []int{1, 1, 1}, []bool{false, false, false}},
		{"TestCase", []int{1, 1, 1, 0, 1}, []bool{false, false, false, false, false}},
		{"TestCase", []int{0, 1, 1, 1, 1, 1}, []bool{true, false, false, false, true, false}},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			got := Solution(c.inputs)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v",
					c.expect, got, c.inputs)
			}
			got = Solution2(c.inputs)
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
