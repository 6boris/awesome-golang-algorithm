package Solution

import (
	"reflect"
	"testing"
)

func TestSolution(t *testing.T) {
	//	测试用例
	cases := []struct {
		name   string
		input1 string
		input2 []string
		expect []int
	}{
		{"TestCacse 1",
			"barfoothefoobarman",
			[]string{"foo", "bar"},
			[]int{0, 9},
		},
		{"TestCacse 2",
			"wordgoodstudentgoodword",
			[]string{"word", "student"},
			[]int{},
		},
	}

	//	开始测试
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := findSubstring(c.input1, c.input2)
			if !reflect.DeepEqual(ret, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v",
					c.expect, ret, c.input1)
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
