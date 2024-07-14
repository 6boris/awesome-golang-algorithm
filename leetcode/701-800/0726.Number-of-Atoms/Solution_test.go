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
		{"TestCase1", "H2O", "H2O"},
		{"TestCase2", "Mg(OH)2", "H2MgO2"},
		{"TestCase3", "K4(ON(SO3)2)2", "K4N2O14S4"},
		{"TestCase4", "H2(H3)2", "H8"},
		{"TestCase5", "H11He49NO35B7N46Li20", "B7H11He49Li20N47O35"},
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
