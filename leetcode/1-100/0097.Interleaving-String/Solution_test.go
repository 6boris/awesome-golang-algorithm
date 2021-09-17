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
		{"TestCase", []string{
			"aabcc", "dbbca", "aadbbcbcac",
		}, true},
		{"TestCase", []string{
			"aabcc", "dbbca", "aadbbbaccc",
		}, false},
	}

	//	开始测试
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := isInterleave(c.inputs[0], c.inputs[1], c.inputs[2])
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v",
					c.expect, got, c.inputs)
			}
		})
	}
}

func TestSolution2(t *testing.T) {
	//	测试用例
	cases := []struct {
		name   string
		inputs []string
		expect bool
	}{
		{"TestCase", []string{
			"aabcc", "dbbca", "aadbbcbcac",
		}, true},
		{"TestCase", []string{
			"aabcc", "dbbca", "aadbbbaccc",
		}, false},
	}

	//	开始测试
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := isInterleave2(c.inputs[0], c.inputs[1], c.inputs[2])
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
