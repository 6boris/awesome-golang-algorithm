package Solution

import (
	"reflect"
	"testing"
)

func TestSolution(t *testing.T) {
	//	测试用例
	cases := []struct {
		name   string
		x int
		expect [][]string
	}{
		{"TestCase", 4, [][]string{{".Q..","...Q","Q...","..Q."},{"..Q.","Q...","...Q",".Q.."}}},
		{"TestCase", 1, [][]string{{"Q"}}},
	}

	//	开始测试
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := Solution(c.x)
			if !reflect.DeepEqual(ret, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v",
					c.expect, ret, c.x)
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
