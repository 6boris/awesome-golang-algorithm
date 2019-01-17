package main

import (
	"github.com/kylesliu/awesome-golang-leetcode/cmd/leetcode"
)

func main() {
	problems := leetcode.GetSotedproblemsInstance()

	//for _, v := range problems {
	//	fmt.Println(v.PathName)
	//}

	//	生成Problem 目录
	//leetcode.MakeDir(problems)

	//leetcode.GetReadmeTemplateBuffer()

	//	GitBook
	leetcode.MakeGitbookSummary(problems)

}
