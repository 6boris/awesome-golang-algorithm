package main

import (
	"fmt"
	"github.com/kylesliu/awesome-golang-leetcode/cmd/leetcode"
)

func main() {
	problems := leetcode.GetProblemsInstance()

	for _, v := range problems {
		if v.PaidOnly {

			fmt.Println(v)
		}
	}
	//	生成Problem 目录
	leetcode.MakeDir(problems)

	leetcode.GetReadmeTemplateBuffer()

}
