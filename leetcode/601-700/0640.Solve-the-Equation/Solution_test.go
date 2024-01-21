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
		expect string
	}{
		{"TestCase1", "x+5-3+x=6+x-2", "x=2"},
		{"TestCase2", "x=x", "Infinite solutions"},
		{"TestCase3", "2x=x", "x=0"},
		{"TestCase4", "2x+3x-6x=x+2", "x=-1"},
		{"TestCase5", "2x+0=0", "x=0"},
		{"TestCase6", "0x=0", "Infinite solutions"},
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

func TestParseExpression(t *testing.T) {
	//	测试用例
	cases := []struct {
		name   string
		inputs string
		a, b   int
	}{
		{"TestCase1", "3x-2x+x-5+3", 2, -2},
		{"TestCase2", "-2x+x-5+3", -1, -2},
		{"TestCase3", "x+5-3+x", 2, 2},
		{"TestCase4", "x", 1, 0},
		{"TestCase5", "6+x-2", 1, 4},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			a, b := parseExpression(c.inputs)
			if a != c.a || b != c.b {
				t.Fatalf("expected: %v %v, but got: %v %v, with inputs: %v",
					c.a, c.b, a, b, c.inputs)
			}
		})
	}
}

func TestGCD(t *testing.T) {
	//	测试用例
	cases := []struct {
		name   string
		a, b   int
		expect int
	}{
		{"TestCase1", 1, 1, 1},
		{"TestCase2", 8, 4, 4},
		{"TestCase3", 12, 2, 2},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			got := gcd640(c.a, c.b)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v %v",
					c.expect, got, c.a, c.b)
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
