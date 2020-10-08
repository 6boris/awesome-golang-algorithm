package Solution

import (
	"strconv"
	"testing"
)

func TestSolution(t *testing.T) {
	//	测试用例
	cases := []struct {
		name   string
		inputs []string
		expect string
	}{
		{"TestCase", []string{"ADOBECODEBANC", "ABC"}, "BANC"},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			//got := minWindow(c.inputs[0], c.inputs[1])
			//if !reflect.DeepEqual(got, c.expect) {
			//	t.Fatalf("expected: %v, but got: %v, with inputs: %v",
			//		c.expect, got, c.inputs)
			//}
		})
	}
}

//	压力测试
func BenchmarkSolution(b *testing.B) {

}

//	使用案列
func ExampleSolution() {

}
