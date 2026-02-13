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
		expect []string
	}{
		{"TestCase1", []string{"pes", "fifa", "gta", "pes(2019)"}, []string{"pes", "fifa", "gta", "pes(2019)"}},
		{"TestCase2", []string{"gta", "gta(1)", "gta", "avalon"}, []string{"gta", "gta(1)", "gta(2)", "avalon"}},
		{"TestCase3", []string{"onepiece", "onepiece(1)", "onepiece(2)", "onepiece(3)", "onepiece"}, []string{"onepiece", "onepiece(1)", "onepiece(2)", "onepiece(3)", "onepiece(4)"}},
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
