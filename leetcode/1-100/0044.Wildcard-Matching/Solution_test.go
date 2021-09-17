package Solution

import (
	"reflect"
	"testing"
)

func TestSolution(t *testing.T) {
	//	测试用例
	cases := []struct {
		name   string
		inputs []string
		expect bool
	}{
		{"TestCase 1", []string{"aa", "a"}, false},
		{"TestCase 2", []string{"aa", "*"}, true},
		{"TestCase 3", []string{"cd", "?a"}, false},
		{"TestCase 4", []string{"adceb", "a*b"}, true},
		{"TestCase 5", []string{"acdcb", "a*c?b"}, false},
	}

	//	开始测试
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := isMatch(c.inputs[0], c.inputs[1])
			if !reflect.DeepEqual(ret, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v",
					c.expect, ret, c.inputs)
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
