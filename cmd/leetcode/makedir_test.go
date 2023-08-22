package leetcode

import (
	"fmt"
	"testing"
)

func TestMakeDir(t *testing.T) {
	problems := GetSortedProblemsInstance()
	fmt.Println(len(problems))

	for _, v := range problems {
		if v.PaidOnly {
			fmt.Println(v)
		}
		fmt.Println(v)
	}

	MakeDir(problems)
}
