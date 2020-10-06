package Solution

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"runtime"
	"testing"
)

//	solution func Info
type SolutionFuncType func(string, int) string

var SolutionFuncList = []SolutionFuncType{
	reverseLeftWords,
	reverseLeftWords2,
}

//	test info struct
type Case struct {
	name   string
	inputs string
	k      int
	expect string
}

// 	test case
var cases = []Case{
	{
		name:   "TestCase 1",
		inputs: "abcdefg",
		k:      2,
		expect: "cdefgab",
	},
	{
		name:   "TestCase 1",
		inputs: "lrloseumgh",
		k:      6,
		expect: "umghlrlose",
	},
}

// TestSolution Example for solution test cases
func TestSolution(t *testing.T) {
	ast := assert.New(t)

	for _, f := range SolutionFuncList {
		for _, c := range cases {
			t.Run(c.name, func(t *testing.T) {
				actual := f(c.inputs, c.k)
				ast.Equal(c.expect, actual,
					"func: %v case: %v ",
					runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name(), c.name)
			})
		}
	}
}
