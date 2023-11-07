package Solution

import (
	"fmt"
	"reflect"
	"runtime"
	"strings"
	"testing"
)

// Solution func Info
type SolutionFuncType func([]string) [][]string

var SolutionFuncList = []SolutionFuncType{
	groupAnagrams,
}

// Test case info struct
type Case struct {
	name   string
	input  []string
	expect [][]string
}

// Test case
var cases = []Case{
	{
		name:   "TestCase 1",
		input:  []string{"eat", "tea", "tan", "ate", "nat", "bat"},
		expect: [][]string{{"bat"}, {"nat", "tan"}, {"ate", "eat", "tea"}},
	},
}

// TestSolution Run test case for all solutions
func TestSolution(t *testing.T) {
	for _, f := range SolutionFuncList {
		funcName := strings.Split(runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name(), ".")[1]
		for _, c := range cases {
			t.Run(fmt.Sprintf("%s %s", funcName, c.name), func(t *testing.T) {
				got := f(c.input)
				fmt.Println(fmt.Printf("want: %+v \n got: %+v", c.expect, got))
				fmt.Println()
			})
		}
	}
}
