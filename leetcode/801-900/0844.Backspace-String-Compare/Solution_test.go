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
		{"TestCase1", "ab#c", "ad#c", true},
		{"TestCase2", "ab##", "c#d#", true},
		{"TestCase3", "a#c", "b", false},
		{"TestCase4", "###ab", "ab", true},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			got := Solution(c.s, c.t)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs-s: %v, inputs-t: %v",
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
