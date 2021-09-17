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
		s, t   string
		expect bool
	}{
		{"TestCase1", "egg", "add", true},
		{"TestCase2", "foo", "bar", false},
		{"TestCase3", "bar", "foo", false},
		{"TestCase4", "paper", "title", true},
		{"TestCase5", "title", "paper", true},
		{"TestCase6", "abcd", "efgh", true},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			got := Solution(c.s, c.t)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v %v",
					c.expect, got, c.s, c.t)
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
