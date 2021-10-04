package Solution

import (
	"reflect"
	"strconv"
	"testing"
)

func TestSolution(t *testing.T) {
	//	测试用例
	cases := []struct {
		name             string
		inputs           [][]int
		sr, sc, newColor int
		expect           [][]int
	}{
		{"TestCase1", [][]int{{1, 1, 1}, {1, 1, 0}, {1, 0, 1}}, 1, 1, 2, [][]int{{2, 2, 2}, {2, 2, 0}, {2, 0, 1}}},
		{"TestCase2", [][]int{{1}}, 0, 0, 2, [][]int{{2}}},
		{"TestCase3", [][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}}, 1, 1, 2, [][]int{{1, 1, 1}, {1, 2, 1}, {1, 1, 1}}},
		{"TestCase4", [][]int{{0, 0, 0}, {0, 0, 0}}, 0, 0, 2, [][]int{{2, 2, 2}, {2, 2, 2}}},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			got := Solution(c.inputs, c.sr, c.sc, c.newColor)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs-image: %v, input-sr: %v, input-sc: %v",
					c.expect, got, c.inputs, c.sr, c.sc)
			}
		})
	}
}
