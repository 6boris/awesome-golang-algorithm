package leetcode

import (
	"fmt"
	"testing"
)

func TestMakeDir(t *testing.T) {
	problems := GetProblemsInstance()

	for _, v := range problems {
		if v.PaidOnly {

			fmt.Println(v)
		}
	}
	//MakeDir(problems)
}
