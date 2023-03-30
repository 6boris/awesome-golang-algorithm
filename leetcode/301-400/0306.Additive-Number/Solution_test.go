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
		{"TestCase1", "112358", true},
		{"TestCase2", "199100199", true},
		{"TestCase3", "000", true},
		{"TestCase4", "101", true},
		{"TestCase5", "199001200", false},
		{"TestCase6", "123", true},
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

func TestAddStr(t *testing.T) {
	cases := []struct {
		s               string
		start, mid, end int
		exp             string
	}{
		{"1234", 0, 1, 3, "64"},
		{"12345", 0, 1, 4, "753"},
		{"99999", 0, 2, 4, "8901"},
	}
	for i, c := range cases {
		t.Run(c.s[c.start:c.mid+1]+"+"+c.s[c.mid+1:c.end+1]+" "+strconv.Itoa(i), func(t *testing.T) {
			got := addStr(c.s, c.start, c.mid, c.end)
			if !reflect.DeepEqual(got, c.exp) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v", c.exp, got, c.s)
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
