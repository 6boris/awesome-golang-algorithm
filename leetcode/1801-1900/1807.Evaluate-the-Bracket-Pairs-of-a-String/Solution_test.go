package Solution

import (
	"reflect"
	"strconv"
	"testing"
)

func TestSolution(t *testing.T) {
	//	测试用例
	cases := []struct {
		name      string
		inputs    string
		knowledge [][]string

		expect string
	}{
		{"TestCase1", "(name)is(age)yearsold", [][]string{
			{"name", "bob"}, {"age", "two"},
		}, "bobistwoyearsold"},
		{"TestCase2", "hi(name)", [][]string{
			{"a", "b"},
		}, "hi?"},
		{"TestCase3", "(a)(a)(a)aaa", [][]string{
			{"a", "yes"},
		}, "yesyesyesaaa"},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			got := Solution(c.inputs, c.knowledge)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v %v",
					c.expect, got, c.inputs, c.knowledge)
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
