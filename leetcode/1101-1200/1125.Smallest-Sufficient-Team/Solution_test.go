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
		inputs []string
		people [][]string
		expect []int
	}{
		{"TestCase1", []string{"java", "nodejs", "reactjs"}, [][]string{
			{"java"}, {"nodejs"}, {"nodejs", "reactjs"},
		}, []int{0, 2}},
		{"TestCase2", []string{"algorithms", "math", "java", "reactjs", "csharp", "aws"}, [][]string{
			{"algorithms", "math", "java"},
			{"algorithms", "math", "reactjs"},
			{"java", "csharp", "aws"},
			{"reactjs", "csharp"},
			{"csharp", "math"},
			{"aws", "java"},
		}, []int{1, 2}},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			got := Solution(c.inputs, c.people)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v %v",
					c.expect, got, c.inputs, c.people)
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
