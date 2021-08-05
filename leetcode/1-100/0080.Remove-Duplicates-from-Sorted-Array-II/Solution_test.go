package Solution

import (
	"reflect"
	"testing"
)

func TestSolution(t *testing.T) {
	//	测试用例
	cases := []struct {
		name   string
		inputs []int
		expect int
	}{
		{"TestCase1", []int{1}, 1},
		{"TestCase2", []int{1, 1, 1, 2, 2, 3}, 5},
		{"TestCase3", []int{0, 0, 1, 1, 1, 1, 2, 3, 3}, 7},
		{"TestCase4", []int{1, 2, 3, 4}, 4},
	}

	//	开始测试
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := Solution(c.inputs)
			if !reflect.DeepEqual(ret, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v",
					c.expect, ret, c.inputs)
			}
			t.Logf("%v", c.inputs[:c.expect])
		})
	}
}

//	压力测试
func BenchmarkSolution(b *testing.B) {

}

//	使用案列
func ExampleSolution() {

}
