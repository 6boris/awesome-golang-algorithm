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
	//leetcode.MakeDirFromTo(problems, 1501, 1600)
	//leetcode.MakeDirFromTo(problems, 1601, 1700)
	//leetcode.MakeDirFromTo(problems, 1701, 1800)
	//leetcode.MakeDirFromTo(problems, 1901, 2000)
	//leetcode.MakeDirFromTo(problems, 2001, 2100)
	//leetcode.MakeDirFromTo(problems, 2101, 2200)
	//leetcode.MakeDirFromTo(problems, 2201, 2300)
	//leetcode.MakeDirFromTo(problems, 2301, 2400)
	//leetcode.MakeDirFromTo(problems, 2401, 2500)
	//leetcode.MakeDirFromTo(problems, 2901, 3000)
	//leetcode.MakeDirFromTo(problems, 3101, 3200)
	//leetcode.MakeDirFromTo(problems, 3201, 3300)
	//leetcode.MakeDirFromTo(problems, 3301, 3400)
	//leetcode.MakeDirFromTo(problems, 3401, 3500)
	//leetcode.MakeDirFromTo(problems, 3601, 3700)
	leetcode.MakeDirFromTo(problems, 3701, 3800)


	// leetcode.MakeDir(problems)

	// leetcode.GetReadmeTemplateBuffer()

	//	GitBook
	// leetcode.MakeGitbookSummary(problems)

	//	sitemap
	// s := sitemap.New(problems)
	// fmt.Println(s)
}
