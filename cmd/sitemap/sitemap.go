package sitemap

import (
	"fmt"
	"github.com/kylesliu/awesome-golang-leetcode/cmd/leetcode"
)

type Config struct {
	SecureKey string
}

type SiteMap struct {
	Url []UrlSet
}

func New(problems []leetcode.Problem) SiteMap {
	fmt.Println(len(problems))
	for i := 0; i < len(problems); i++ {
		fmt.Println(problems[i].DirPath)
	}
	return SiteMap{}
}

type UrlSet struct {
	Loc     string
	Lastmod string
}
