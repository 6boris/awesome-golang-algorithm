package main

import (
	"awesome-golang-algorithm/cmd/leetcode"
	"fmt"
)

func main() {
	problems := leetcode.GetSortedProblemsInstance()
	//	fmt.Println(v)
	for _, v := range problems {
		fmt.Println(v)
	}
	//problems := leetcode.GetProblemsJson()
	//fmt.Println(problems)
	//for _, v := range problems {
	//	fmt.Println(v.Stat.QuestionID, v.Stat.QuestionTitle, v.PathName,v.Stat.QuestionTitleSlup)
	//err := leetcode.UrlPath("https://leetcode.com/problems/" + v.Stat.QuestionTitleSlup )
	//if err != nil {
	//	fmt.Println(v.Stat.QuestionID, v.Stat.QuestionTitle, v.PathName,v.Stat.QuestionTitleSlup)
	//	fmt.Println(err.Error())
	//}
	//}

	//	生成Problem 目录

	// leetcode.MakeDir(problems)

	// leetcode.GetReadmeTemplateBuffer()

	//	GitBook
	// leetcode.MakeGitbookSummary(problems)

	//	sitemap
	// s := sitemap.New(problems)
	// fmt.Println(s)
}
