package Solution

import (
	"fmt"
	"reflect"
	"runtime"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

//	solution func Info
type SolutionFuncType func(string, string) bool

var SolutionFuncList = []SolutionFuncType{
	// isAnagram_1,
	// isAnagram_2,
	isAnagram_3,
}

//	test info struct
type Case struct {
	name   string
	s      string
	t      string
	expect bool
}

// 	test case
var cases = []Case{
	{name: "TestCase 1", s: "anagram", t: "nagaram", expect: true},
	{name: "TestCase 2", s: "rat", t: "car", expect: false},
	{name: "TestCase 3", s: "ab", t: "a", expect: false},
	{name: "TestCase 4", s: "aa", t: "bb", expect: false},
	{name: "TestCase 5", s: "xaaddy", t: "xbbccy", expect: false},
}

// TestSolution Example for solution test cases
func TestSolution(t *testing.T) {
	ast := assert.New(t)

	for _, f := range SolutionFuncList {
		funcName := strings.Split(runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name(), ".")[1]
		for _, c := range cases {
			t.Run(fmt.Sprintf("%s %s", funcName, c.name), func(t *testing.T) {
				got := f(c.s, c.t)
				ast.Equal(c.expect, got,
					"func: %v case: %v ", funcName, c.name)
			})
		}
	}
}
