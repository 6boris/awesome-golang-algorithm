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
		segs   []seg
		res    []bool
		expect bool
	}{
		{"TestCase1", []seg{{2, 6}, {12, 15}, {8, 9}, {9, 10}}, []bool{true, true, true, true}, true},
		{"TestCase2", []seg{{10, 20}, {15, 25}, {20, 30}}, []bool{true, false, true}, true},
		{"TestCase3", []seg{{37, 50}, {33, 50}, {4, 17}, {35, 48}, {8, 25}}, []bool{true, false, true, false, false}, true},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			got := Solution(c.segs, c.res)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with input-segs: %v, input-res: %v",
					c.expect, got, c.segs, c.res)
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
