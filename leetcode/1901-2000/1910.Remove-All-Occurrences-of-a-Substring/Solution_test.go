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
		s, p   string
		expect string
	}{
		{"TestCase1", "daabcbaabcbc", "abc", "dab"},
		{"TestCase2", "axxxxyyyyb", "xy", "ab"},
		{"TestCase3", "eemckxmckx", "emckx", ""},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			got := Solution(c.s, c.p)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v %v",
					c.expect, got, c.s, c.p)
			}
		})
	}
}

func TestSolution1(t *testing.T) {
	//	测试用例
	cases := []struct {
		name   string
		s, p   string
		expect string
	}{
		{"TestCase1", "daabcbaabcbc", "abc", "dab"},
		{"TestCase2", "axxxxyyyyb", "xy", "ab"},
		{"TestCase3", "eemckxmckx", "emckx", ""},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			got := Solution1(c.s, c.p)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v %v",
					c.expect, got, c.s, c.p)
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
