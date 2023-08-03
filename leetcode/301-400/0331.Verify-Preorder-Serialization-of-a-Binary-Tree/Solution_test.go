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
		inputs string
		expect bool
	}{
		{"TestCase1", "9,3,4,#,#,1,#,#,2,#,6,#,#", true},
		{"TestCase2", "1,#", false},
		{"TestCase3", "9,3,#,1", false},
		{"TestCase4", "9,#,93,#,9,9,#,#,#", true},
		{"TestCase5", "9,9,91,#,#,9,#,49,#,#,#", true},
		{"TestCase6", "8,#,5,#,2,5,#,7,9,#,8,#,#,#,#", true},
		{"TestCase7", "1,#,#,#,#", false},
		{"TestCase8", "1,#,#", true},
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

// 压力测试
func BenchmarkSolution(b *testing.B) {
}

// 使用案列
func ExampleSolution() {
}
