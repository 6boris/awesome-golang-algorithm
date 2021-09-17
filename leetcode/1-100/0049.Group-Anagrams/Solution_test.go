package Solution

import (
	"fmt"
	"testing"
)

func TestSolution(t *testing.T) {
	//	测试用例
	cases := []struct {
		name   string
		inputs []string
		expect [][]string
	}{
		{
			"1 test 1",
			[]string{"eat", "tea", "tan", "ate", "nat", "bat"},
			[][]string{
				{"eat", "tea", "ate"},
				{"tan", "nat"},
				{"bat"},
			},
		},
	}

	//	开始测试
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := groupAnagrams(c.inputs)
			fmt.Println(got)
			//if !reflect.DeepEqual(ret, c.expect) {
			//	t.Fatalf("expected: %v, but got: %v, with inputs: %v",
			//		c.expect, ret, c.inputs)
			//}
		})
	}
}
